package Process

import (
	"../../Common/Message"
	"../Utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	//暂时不需要字段..
}



//完成登录功能的函数,将登陆信息发送给服务器，返回err，可以看到更多登录失败的信息
//发送消息 1.发送信息的长度 2.发送信息
func (this *UserProcess)Login(userId int,userPwd string)(error){

	//fmt.Printf(" userId = %d userPwd=%s\n", userId, userPwd)
	//1.链接到服务器
	conn,err:=net.Dial("tcp","localhost:8889")
	//链接必须延迟关闭
	defer conn.Close()
	if(err!=nil){
		fmt.Println("net.Dial err=",err)
		return err;
	}

	//2.准备通过conn连接发送登陆消息给客户端
	var message Message.Message
	message.Type =Message.LoginMesType
	//3.创建一个message的数据部分，为LoginMes结构体
	var loginMes Message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//message.Data = loginMes  错误message.Data是string类型，需要将loginMes序列化
	//4.将loginMes序列化,返回一个byte切片
	fmt.Println("loginMes",loginMes)
	dataLoginMes,err:=json.Marshal(loginMes)
	if(err!=nil){
		fmt.Println("json.Marshal(loginMes) err=",err)
		return err;
	}

	//5.将dataloginMes赋值给message.Data
	message.Data = string(dataLoginMes)

	//6.将message序列化
	dataMessage,err:=json.Marshal(message)
	if(err!=nil){
		fmt.Println("json.Marshal(message) err=",err)
		return err;
	}

	//7.调用Util.WtitwPkg函数
	//发送message序列化后的长度与数据本身值给服务器
	transfer:=Utils.Transfer{Conn: conn}
	errr:=transfer.WtitePkg(dataMessage)
	if(errr!=nil){
		fmt.Println("conn.Write(dataMessage) err")
		return errr;
	}
	fmt.Println("客户端，发送dataMessage信息完成")

	//1.处理服务器返回的消息
	message,err =transfer.ReadPkg()
	if(err!=nil){
		fmt.Println("客户端：Utils.ReadPkg(conn) err")
		return errr;
	}
	//2.将message的data部分反序列成RespondLoginmes
	var respondLoginMes Message.RespondLoginMes
	err = json.Unmarshal([]byte(message.Data),&respondLoginMes)
	if(err!=nil){
		fmt.Println("客户端：json.Unmarshal([]byte(message.Data),respondLoginMes) err")
		return errr;
	}
	fmt.Println("客户端接受消息完毕，消息：",respondLoginMes)

	//3.对RespondLoginmes的状态码进行校验
	if respondLoginMes.StatusCode == 200 {
		//fmt.Println("经服务器校验登陆信息：登陆成功")

		//遍历respondLoginMes的Ids字段，显示当前在线用户
		for index,userId:=range respondLoginMes.UsersIds{
			fmt.Printf("第%v个在线用户的ID是，%v\n",index,userId)
		}

		//该协程保持和服务器端的通讯.如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端.
		go serverProcessMes(conn)

		//登陆成功，调用server中函数，显示我们的登录成功的菜单
		for{
			//调用server中函数，显示我们的登录成功的菜单
			ShowMenu();
		}
		return nil
	}else if respondLoginMes.StatusCode == 500 {
		//fmt.Println("经服务器校验登陆信息：登陆失败，请注册")
		return errors.New("经服务器校验登陆信息：登陆失败，请注册")
	}else if respondLoginMes.StatusCode ==403{
		//fmt.Println("错误的状态码")
		return errors.New("登陆消息返回：密码错误")
	}

	return nil
}


//完成注册功能的函数,将注册信息发送给服务器，返回err，可以看到更多登录失败的信息
//发送消息 1.发送信息的长度 2.发送信息
func (this *UserProcess)Register(userId int,userPwd string,userName string)(error){

	//fmt.Printf(" userId = %d userPwd=%s\n", userId, userPwd)
	//1.链接到服务器
	conn,err:=net.Dial("tcp","localhost:8889")
	//链接必须延迟关闭
	defer conn.Close()
	if(err!=nil){
		fmt.Println("net.Dial err=",err)
		return err;
	}

	//2.准备通过conn连接发送登陆消息给客户端
	var message Message.Message
	message.Type =Message.RegisterMesType
	//3.创建一个message的数据部分，为RegisterMes结构体
	var registerMes Message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName


	//message.Data = registerMes  错误message.Data是string类型，需要将registerMes序列化
	//4.将registerMes序列化,返回一个byte切片
	fmt.Println("registerMes:",registerMes)
	MarshalDataRegisterMes,err:=json.Marshal(registerMes)
	if(err!=nil){
		fmt.Println("客户端：json.Marshal(registerMes) err=",err)
		return err;
	}

	//5.将dataloginMes赋值给message.Data
	message.Data = string(MarshalDataRegisterMes)

	//6.将message序列化
	dataMessage,err:=json.Marshal(message)
	if(err!=nil){
		fmt.Println("客户端：json.Marshal(message) err=",err)
		return err;
	}

	//7.调用Util.WtitwPkg函数
	//发送message序列化后的长度与数据本身值给服务器
	transfer:=Utils.Transfer{Conn: conn}
	errr:=transfer.WtitePkg(dataMessage)
	if(errr!=nil){
		fmt.Println("客户端：conn.Write(dataMessage) err")
		return errr;
	}
	fmt.Println("客户端，发送dataMessage信息完成")

	//1.处理服务器返回的消息
	message,err =transfer.ReadPkg()
	if(err!=nil){
		fmt.Println("客户端：Utils.ReadPkg(conn) err")
		return errr;
	}
	//2.将message的data部分反序列成RespondLoginmes
	var respondRegisterMes Message.RespondRegisterMes
	err = json.Unmarshal([]byte(message.Data),&respondRegisterMes)
	if(err!=nil){
		fmt.Println("客户端：json.Unmarshal([]byte(message.Data),&respondRegisterMes) err")
		return errr;
	}
	fmt.Println("客户端接受消息完毕，消息：",respondRegisterMes)

	//3.对RespondLoginmes的状态码进行校验
	if respondRegisterMes.StatusCode == 200 {
		fmt.Println("经服务器校验注册信息：注册成功，请登录")
		//注册成功，重回开始界面，让用户进行登录

	}else if respondRegisterMes.StatusCode == 400 {
		fmt.Println("经服务器校验注册信息：注册失败，用户已注册")
		return errors.New("经服务器校验注册信息：注册失败，用户已注册")
	}else if respondRegisterMes.StatusCode == 403{
		fmt.Println("经服务器校验注册信息：注册失败，未知错误")
		return errors.New("经服务器校验注册信息：注册失败，未知错误")
	}
	return nil
}