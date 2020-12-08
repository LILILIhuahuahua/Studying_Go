package main

import (
	"../Process"
	"fmt"
	"os"
)
//定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string
var userName string



func main(){
	//接受用户选择
	var key int;
	var loop = true;

	for loop{
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		//一行为一个输入
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1 :
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			//1. 创建一个UserProcess的实例
			userProcess := &Process.UserProcess{}
			err :=userProcess.Login(userId, userPwd)
			if(err!=nil){
				fmt.Println("登陆失败：err",err)
			}else{
				fmt.Println("登陆成功(服务器已校验id与pwd)")
			}
		case 2 :
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			//2. 调用UserProcess，完成注册的请求、
			userProcess := &Process.UserProcess{}
			userProcess.Register(userId, userPwd, userName)
		case 3 :
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default :
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}