package process

import (
	"fmt"
	"sync"
)

// UserManager实例在服务端有且仅有一个
// 所以直接定义为全局变量，供需要的地方使用
var (
	MyUserManager *UserManager
)

// UserManager实现对在线用户的管理
type UserManager struct {
	// map是非线程安全的，并发读写操作时 可能会出现错误；可用并发安全的sync.Map代替
	// https://www.cnblogs.com/qcrao-2018/p/12833787.html
	onlineUsers map[int]*UserProcess
	// 读写锁，保证并发中对onlineUsers读写操作的安全
	rwMutex sync.RWMutex
}

func init() {
	MyUserManager = &UserManager{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 添加在线用户 (同时也支持修改)
func (this *UserManager) AddOnlineUser(up *UserProcess) {
	this.rwMutex.Lock()
	this.onlineUsers[up.UserId] = up
	this.rwMutex.Unlock()
}

// 通过用户Id移除指定的在线用户
func (this *UserManager) DelOnlineUser(userId int) {
	//this.rwMutex.Lock()
	delete(this.onlineUsers, userId)
	//this.rwMutex.Unlock()
}
// 通过ConnAddr移除指定的在线用户
func (this *UserManager) DelOnlineUserByAddr(connAddr string) {
	// 遍历在线用户列表，通过连接地址寻找到对应的用户Id
	this.rwMutex.RLock()
	for uid, up := range this.onlineUsers {
		if up.ConnAddr == connAddr {
			this.DelOnlineUser(uid)
			break
		}
	}
	this.rwMutex.RUnlock()
}

// 获取当前所有在线的用户
func (this *UserManager) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 获取指定的用户的UserProcess
func (this *UserManager) GetOnlineUserById(userId int) (*UserProcess, error) {

	this.rwMutex.RLock()
	up, ok := this.onlineUsers[userId]
	this.rwMutex.RUnlock()
	if ok {
		return up, nil
	}
	return nil, fmt.Errorf("用户[%d]不存在", userId)
}

// 获取指定的用户的UserProcess (通过连接地址)
func (this *UserManager) GetOnlineUserByAddr(connAddr string) (up *UserProcess, err error) {
	// 遍历在线用户列表，通过连接地址寻找到对应的用户
	this.rwMutex.RLock()
	for _, _up := range this.onlineUsers {
		if _up.ConnAddr == connAddr {
			up = _up
			break
		}
	}
	this.rwMutex.RUnlock()
	if up == nil {
		err = fmt.Errorf("用户不存在")
	}
	return up, err
}