package Process
// 管理一个在线用户的切片  map【用户ID】【UserProcess】

import (
	"fmt"
)

//因为UserMgr 实例在服务器端有且只有一个
//因为在很多的地方，都会使用到，因此，我们
//将其定义为全局变量
var (
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int]*UserProcess //map【用户ID】【UserProcess指针】
}
//完成对userMgr初始化工作 （给onlineUsers切片分配空间）
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}


//对onlineUsers（map）的增删改查的操作
//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(userProcess *UserProcess) {
	this.onlineUsers[userProcess.UserId] = userProcess
}
//当用户离线，在onlineUsers中删除该用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值  (便于服务器连接两个客户端)
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明，你要查找的这个用户，当前不在线。
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}