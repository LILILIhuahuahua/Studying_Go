package main

import (
	"errors"
	"fmt"
	"net"
	"../Process"
	"../../Common/Message"
	"../Utils"
)

//先创建一个Processor 的结构体体,负责消息的接收消息与消息转发
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，调用不同函数
func (this *Processor)SwitchServerFucntion(message *Message.Message)(err error){

	switch message.Type {
	//Processor根据所需不同的服务，去创建不同Process实例去执行服务
	case Message.LoginMesType:
		//1.处理登陆逻辑
		//1.1创建一个UserProcess实例
		userProcess := &Process.UserProcess{
			Conn : this.Conn,
		}
		//fmt.Printf("switchServerFucntion:  message.data=%v  message.data类型：%p\n", message.Data,message.Data)
		fmt.Println("消息种类：登陆")
		userProcess.LoginServerProcess(message)
	case Message.RegisterMesType:
		//2.处理注册逻辑
		//2.1创建一个UserProcess实例
		userProcess := &Process.UserProcess{
			Conn : this.Conn,
		}
		//fmt.Printf("switchServerFucntion:  message.data=%v  message.data类型：%p\n", message.Data,message.Data)
		fmt.Println("消息种类：注册")
		userProcess.RegisterServerProcess(message)
	default:
		fmt.Println("消息种类不存在，无法处理")
		return errors.New("消息种类不存在，无法处理")

	}
	return nil
}

//main开启一个协程后，通过协程调用process2循环接受客户端的消息
func (this *Processor) GetMesFromClient()(err error){
	//要读取或传输数据，创建一个Transfer实例
	transfer:=Utils.Transfer{Conn: this.Conn}

	//一直读取客户端发送的数据
	for{
		//读取包的代码，封装成readPkg(),返回Message，err
		fmt.Println("----服务器又一次读取信息------")

		//要读取或传输数据，创建一个Transfer实例
		//transfer:=Utils.Transfer{Conn: this.Conn}
		messaage,err:=transfer.ReadPkg()
		if(err!=nil){
			fmt.Println("服务器：readPkg(conn) err")
			return err
		}
		fmt.Println("mess= ",messaage)

		//读取的消息，给switchServerFucntion进行下一步的处理
		err=this.SwitchServerFucntion(&messaage)
		if(err!=nil){
			fmt.Println("readPkg(conn) err",)
			return err
		}
	}
}