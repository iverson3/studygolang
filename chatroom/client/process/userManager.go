package process

import (
	"fmt"
	model2 "studygolang/chatroom/client/model"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
)

// 客户端维护的在线好友列表 (map是非线程安全的)
var onlineUserList = make(map[int]*model.User, 10)
// 全局的客户端当前用户model
var CurUser model2.CurUser

func ShowOnlineUserList() {
	fmt.Println("在线好友列表：")
	fmt.Println("=====================")
	for uid, user := range onlineUserList {
		fmt.Printf("[%d] %s - (%s) \n", uid, user.UserName, user.UserStatus)
	}
	fmt.Println("=====================")
}

func UpdateOnlineUserList(notifyMess *common.NotifyUserStatusMes) {
	user, ok := onlineUserList[notifyMess.UserId]
	if !ok {
		user = &model.User{
			UserId:     notifyMess.UserId,
			UserName:   notifyMess.UserName,
		}
	}
	user.UserStatus = notifyMess.Status
	onlineUserList[notifyMess.UserId] = user

	ShowOnlineUserList()
}

// 判断用户列表中是否存在指定的用户
func ExistUserInUserList(userId int) bool {
	_, ok := onlineUserList[userId]
	return ok
}

// 从在线好友列表中移除指定的用户
func RemoveUserFromOnlineUserList(userId int) (bool, error) {
	_, ok := onlineUserList[userId]
	if ok {
		delete(onlineUserList, userId)
	}
	return true, nil
}

// 从在线好用列表中获取指定的用户
func GetUserFromOnlineUserList(userId int) *model.User {
	user, ok := onlineUserList[userId]
	if ok {
		return user
	}
	return nil
}