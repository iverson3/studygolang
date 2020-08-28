package process

import (
	"fmt"
)

// UserManager实例在服务端有且仅有一个
// 所以直接定义为全局变量，供需要的地方使用
var (
	userManager *UserManager
)

// UserManager实现对在线用户的管理
type UserManager struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userManager = &UserManager{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 添加在线用户 (同时也支持修改)
func (this *UserManager) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 移除指定的在线用户
func (this *UserManager) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 获取当前所有在线的用户
func (this *UserManager) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 获取指定的用户的UserProcess
func (this *UserManager) GetOnlineUserById(userId int) (*UserProcess, error) {
	up, ok := this.onlineUsers[userId]
	if ok {
		return up, nil
	}
	return nil, fmt.Errorf("用户[%d]不存在", userId)
}
