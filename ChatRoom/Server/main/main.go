package main

import (
	"fmt"
	"net"
	"time"
	"../Model"
)
func init() {
		//当服务器启动时，我们就去初始化我们的redis的连接池
	    //pool在redis包中被创建和初始化
		initPool("localhost:6379", 16, 0, 300 * time.Second)
		initUserDao()
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool 本身就是一个全局的变量
	//这里需要注意一个初始化顺序问题
	//1.initPool, 2.initUserDao

	Model.MyUserDao = Model.NewUserDao(pool)
}



//由协程保持和客户端通信
func process(conn net.Conn)(err error){
	//开始一个协程后，当协程函数执行结束要将协程关闭（连接需要记得延迟关闭）
	defer conn.Close()

	//这里调用总控, 创建一个
	processor := &Processor{Conn: conn}

	//main开启一个协程后，通过协程总控processor调用循环接受客户端的消息
	err = processor.GetMesFromClient()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误err")
		return
	}

	fmt.Println("process结束")
	return nil
}

func main(){

	//提示信息
	fmt.Println("服务器在8889端口监听")
	listen,err:=net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close() //连接需要记得延迟关闭
	if(err!=nil){
		fmt.Println("net.listen err=",err)
	}

	//一旦监听成功，就等待客户端来连接服务器
	for{
		conn,err:=listen.Accept()
		if(err!=nil){
			fmt.Println("listen.Accept err=",err)
		}

		//一旦连接成功，开始一个协程与客户端保持通信
		go process(conn)
	}
}
