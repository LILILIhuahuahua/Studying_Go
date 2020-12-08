package Model

import (
	"encoding/json"
	"fmt"
	"github.com/redigo/redis"
	 . "../../Common/User"
)

//定义一个UserDao 结构体体,
//UserDao负责去访问数据库，完成对User 结构体的各种操作.

type UserDao struct {
	pool  *redis.Pool
}
//我们在服务器启动后，就初始化一个userDao实例，
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//使用工厂模式，创建一个userDao实例
func NewUserDao(pool  *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

//思考一下在UserDao 应该提供哪些方法给我们
//1. 根据用户id 返回 一个User实例+err
                                //redis.Conn  从redis线程池中获取的一个连接
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {
		//错误!
		//表示在 users 哈希中，没有找到对应id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//这里我们需要把res 反序列化成User实例    (反序列化传指针)
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("UserDao:json.Unmarshal err=", err)
		return
	}
	fmt.Printf("根据id=%v取出user:%v",id,user)
	return user,nil
}


//完成登录的校验 Login
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao 的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("userDao文件:this.getUserById(conn, userId) err")
		return &User{},err
	}
	//这时证明这个用户是获取到.
	//在判断用户密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return &User{},err
	}
	return user,nil
}


//完成注册的校验 Register
//1. Register 检查用户名在数据库中是否已经存在
//2. 如果用户名UserId在数据库中未存在，则返回一个空user实例
//3. 如果用户名UserId在数据库中已存在，则返回对应的错误信息
func (this *UserDao) Register(user User) (err error) {

	//先从UserDao 的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//err!=nil,用户名UserId在数据库中未存在
	fmt.Println("userDao文件:用户名在数据库中不存在，可以注册")
	data,err:=json.Marshal(user)
	if err != nil {
		return
	}
	//将userId与序列化之后的user存入数据库之中
	//入库                                         (记得将data这个切片类型转成string类型)
	_,err =conn.Do("HSet", "user", user.UserId,string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	fmt.Println("userDao.Register存入数据库数据为：",string(data))
	return nil

}