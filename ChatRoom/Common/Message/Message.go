package Message

import (
	"../User"
)

//json是表示通过json传送时，key的值为这里定义的

const(
	LoginMesType = "LoginMes"
	RespondLoginMesType = "RespondLoginMesType"
	RegisterMesType = "RegisterMes"
	RespondRegisterMesType= "RespondRegisterMes"

)

type Message struct {
	Type string `json:"type"`  //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息..后面需要再增加

//登陆消息结构
type LoginMes struct {
	UserId int `json:"userId"` //用户id
	UserPwd string `json:"userPwd"` //用户密码
	UserName string `json:"userName"` //用户名
}
//登陆响应消息结构
//返回登陆消息时，要带上所有已经登陆用户id的切片，用于客户端显示用户列表
type RespondLoginMes struct {
	StatusCode int    `json:"statusCode"` // 返回状态码 500 表示该用户未注册 200表示登录成功，403表未直错误
	UsersIds   []int  // 增加字段，保存用户id的切片
	Error      string `json:"error"` // 返回错误信息
}


//注册消息结构
type RegisterMes struct {
	User User.User  //类型就是User结构体

}

//注册响应消息结构
type RespondRegisterMes struct {
	StatusCode int  `json:"statusCode"` // 返回状态码  400表示该用户已注册 200表示注册成功，403表未直错误
	Error string `json:"error"` // 返回错误信息
}