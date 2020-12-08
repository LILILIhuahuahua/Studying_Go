package Process

import (
	"../../Common/Message"
	"../Utils"
	"encoding/json"
	"fmt"
	"net"
	"../Model"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}




//编写一个LoginServerProcess函数
//功能：在服务器端，实现登陆功能函数，由switchServerFucntion调用
func (this *UserProcess)LoginServerProcess(message *Message.Message)(err error){
	//1.从message中获取message.data部分，直接反序列化为data
	var loginMes Message.LoginMes
	err = json.Unmarshal([]byte(message.Data),&loginMes)
	if(err!=nil){
		fmt.Println("服务器：json.Unmarshal([]byte(message.Data),loginMes) err=")
	}

	//服务器对客户端发来的消息处理后，需要响应
	//2.声明一个respondMes，用于对客户端登陆消息的响应
	var respondMes Message.Message
	respondMes.Type  = Message.RespondLoginMesType
	var respondLoginMes Message.RespondLoginMes //respondLoginMes由下面合法性判断结果决定

	//3.将收到的消息中UserId，传给UserDao，让他去数据库中校验，返回一个User对象与err错误
	user,err:=Model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if(err!=nil){
		//登陆信息不合法
		fmt.Println("数据库校验：登陆信息不合法")
		if err==Model.ERROR_USER_NOTEXISTS{
			respondLoginMes.StatusCode =500 //状态码500表用户未注册
			respondLoginMes.Error = "该用户未注册，请注册在使用"
		}else if err==Model.ERROR_USER_PWD{
			respondLoginMes.StatusCode =403 //状态码403表用户密码错误
			respondLoginMes.Error = "输入密码错误"
		}else{
			respondLoginMes.StatusCode =403 //状态码403表用户密码错误
			respondLoginMes.Error = "未错误"
		}

	}else {
		//登陆信息合法
		fmt.Println("数据库校验：登陆信息合法")
		fmt.Println("数据库中获取的登陆的user全部信息：",user)
		respondLoginMes.StatusCode =200 //状态码200表登陆成功
		respondLoginMes.Error = ""


		//返回消息respondLoginMes的userIds字段：将所有已经登陆的用户userId添加进登陆

		//2.再将登陆成功的用户自己的UserId和自己的userProcess 添加入userManager文件的onlineUsers切片中
		this.UserId = user.UserId
		userMgr.AddOnlineUser(this)
		//1.遍历useMgr的onlineUsers切片
		for userId,_:=range userMgr.onlineUsers{
			respondLoginMes.UsersIds = append(respondLoginMes.UsersIds,userId)
		}

	}



	//4.将respondLoginMes序列化
	datarespondLoginMes,err:=json.Marshal(respondLoginMes)
	if(err!=nil){
		fmt.Println("json.Marshal(respondLoginMes) err=",err)
		return err
	}
	//5.将序列化后的值给respondMes.data
	respondMes.Data = string(datarespondLoginMes)
	//6.序列化整个respondMes
	data,err :=json.Marshal(respondMes)
	if(err!=nil){
		fmt.Println("服务器：json.Marshal(respondMes) err=",err)
		return err
	}

	//7.调用wtitePkg，将respondMes序列化的切片data发送给客户端
	tranfer:= &Utils.Transfer{Conn: this.Conn}
	err= tranfer.WtitePkg(data)
	if(err!=nil){
		fmt.Println("服务器：wtitePkg(conn,data) err")
		return err
	}
	return nil
}

//编写一个RegisterServerProcess函数
//功能：在服务器端，实现注册功能函数，由switchServerFucntion调用
func (this *UserProcess)RegisterServerProcess(message *Message.Message)(err error){
	//1.从message中获取message.data部分，直接反序列化为data
	var registerMes Message.RegisterMes
	err = json.Unmarshal([]byte(message.Data),&registerMes)
	if(err!=nil){
		fmt.Println("服务器：json.Unmarshal([]byte(message.Data),&registerMes)")
	}

	//服务器对客户端发来的消息处理后，需要响应
	//2.声明一个respondMes，用于对客户端登陆消息的响应
	var respondMes Message.Message
	respondMes.Type  = Message.RespondRegisterMesType
	var respondRegisterMes Message.RespondRegisterMes //respondRegisterMes由下面合法性判断结果决定

	//3.将收到的消息中User对象，传UserDao，让他去数据库进行注册校验，返回一个error错误
	err =Model.MyUserDao.Register(registerMes.User)
	if(err!=nil){
		//登陆信息不合法
		fmt.Println("数据库校验：注册信息不合法")
		if err==Model.ERROR_USER_EXISTS{
			respondRegisterMes.StatusCode =400 //状态码400表用户已注册
			respondRegisterMes.Error = "该用户已注册，请务重复注册"
		}else {
			respondRegisterMes.StatusCode =403 //状态码403表未知错误
			respondRegisterMes.Error = "未知错误"
		}
	}else {
		//注册成功
		fmt.Println("数据库校验：注册成功")
		respondRegisterMes.StatusCode =200 //状态码200表登陆成功
		respondRegisterMes.Error = ""
	}


	//4.将respondLoginMes序列化
	datarespondRegisterMes,err:=json.Marshal(respondRegisterMes)
	if(err!=nil){
		fmt.Println("服务器：json.Marshal(respondRegisterMes) err=",err)
		return err
	}
	//5.将序列化后的值给respondMes.data
	respondMes.Data = string(datarespondRegisterMes)
	//6.序列化整个respondMes
	data,err :=json.Marshal(respondMes)
	if(err!=nil){
		fmt.Println("服务器.RegisterServerProcess：json.Marshal(respondMes) err=",err)
		return err
	}

	//7.调用wtitePkg，将respondMes序列化的切片data发送给客户端
	tranfer:= &Utils.Transfer{Conn: this.Conn}
	err= tranfer.WtitePkg(data)
	if(err!=nil){
		fmt.Println("服务器.RegisterServerProcess：wtitePkg(conn,data) err")
		return err
	}

	return nil
}