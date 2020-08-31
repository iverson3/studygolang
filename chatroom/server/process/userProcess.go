package process

import (
	"encoding/json"
	"fmt"
	"net"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
	"studygolang/chatroom/server/utils"
)

type UserProcess struct {
	UserId int
	UserName string
	// 当前用户客户端与服务端所建立的连接
	Conn net.Conn
	// 客户端连接地址 (ip+port)
	ConnAddr string
}

// 根据错误类型判断并返回对应的错误码和错误信息
func judgeError(err error) (code int, errInfo string) {
	if err != nil {
		errInfo = err.Error()
		// 判断不同的错误类型
		if err == model.ERROR_USER_NOTEXISTS {
			code = 500
		} else if err == model.ERROR_USER_PWD {
			code = 403
		} else if err == model.ERROR_USER_EXISTS {
			code = 405
		} else if err == model.ERROR_REDIS_DO_FAILED {
			code = 506
		} else if err == model.ERROR_JSON_MARSHAL {
			code = 501
		} else {
			code = 505
			errInfo = model.ERROR_UNKNOWN.Error()
		}
	} else {
		code = 200
	}
	return
}

func (this *UserProcess) ServerProcessLogin(mess *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mess.Data), &loginMes)
	if err != nil {
		return
	}

	var resMess common.Message
	resMess.Type = common.LoginResMesType

	var loginResMess common.LoginResMes

	// 登录验证
	curUser, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	loginResMess.Code, loginResMess.Error = judgeError(err)

	// 登录成功时，服务端相应的处理
	if loginResMess.Code == 200 {
		fmt.Printf("用户[%s]登录成功；客户端ip地址：%s", curUser.UserName, this.Conn.RemoteAddr().String())

		this.UserId = curUser.UserId
		this.UserName = curUser.UserName
		this.ConnAddr = this.Conn.RemoteAddr().String()
		MyUserManager.AddOnlineUser(this)
		go this.NotifyOtherOnlineUsers(curUser)

		// 构建在线用户列表数据，返回给客户端
		userList := make(map[int]string)
		users := MyUserManager.GetAllOnlineUser()

		for _, user := range users {
			userList[user.UserId] = user.UserName
		}
		loginResMess.OnlineUserList = userList
		loginResMess.UserName = curUser.UserName
	}

	bytes, err := json.Marshal(loginResMess)
	if err != nil {
		return
	}
	resMess.Data = string(bytes)

	data, err := json.Marshal(resMess)
	if err != nil {
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	return tf.WritePkg(data)
}

func (this *UserProcess) ServerProcessRegister(mess *common.Message) (err error) {
	var regMes common.RegisterMes
	err = json.Unmarshal([]byte(mess.Data), &regMes)
	if err != nil {
		return
	}

	var resMess common.Message
	resMess.Type = common.RegisterResMesType

	var regResMess common.RegisterResMes

	// 注册实现
	err = model.MyUserDao.Register(&regMes.User)
	regResMess.Code, regResMess.Error = judgeError(err)

	bytes, err := json.Marshal(regResMess)
	if err != nil {
		return
	}
	resMess.Data = string(bytes)

	data, err := json.Marshal(resMess)
	if err != nil {
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	return tf.WritePkg(data)
}

// 通知其他用户 当前上线用户的在线状态
// 当前用户(userId)通知其他所有在线用户-自己上线了
func (this *UserProcess) NotifyOtherOnlineUsers(user *model.User) {
	data, err := assembleNotifyUserStatusData(user, 1)
	if err != nil {
		return
	}
	_ = this.NotifyOnlineUserStatusChange(user, data)
}

// 通知其他用户 当前用户的下线状态
// 当前用户(userId)通知其他所有在线用户-自己下线了
func (this *UserProcess) NotifyOtherOfflineUsers(user *model.User) {
	data, err := assembleNotifyUserStatusData(user, 2)
	if err != nil {
		return
	}
	_ = this.NotifyOnlineUserStatusChange(user, data)
}


// 组装用户上下线状态消息数据
func assembleNotifyUserStatusData(user *model.User, status int) (data []byte, err error) {
	var mess common.Message
	mess.Type = common.NotifyUserStatusMesType

	var notifyMess common.NotifyUserStatusMes
	notifyMess.UserId = user.UserId
	notifyMess.UserName = user.UserName
	switch status {
	case 1:
		notifyMess.Status = common.UserOnline
	case 2:
		notifyMess.Status = common.UserOffline
	case 3:
		notifyMess.Status = common.UserBusy
	default:
		// 未知状态
	}

	bytes, err := json.Marshal(notifyMess)
	if err != nil {
		return
	}
	mess.Data = string(bytes)

	return json.Marshal(mess)
}

// 将用户状态变化的通知信息数据广播给所有已在线的用户
func (this *UserProcess) NotifyOnlineUserStatusChange(user *model.User, data []byte) (err error) {
	for uid, up := range MyUserManager.onlineUsers {
		if uid != user.UserId {
			tf := &utils.Transfer{Conn: up.Conn}
			err = tf.WritePkg(data)
			if err != nil {
				fmt.Println("通知用户上线失败，失败的用户: ", uid)
			}
		}
	}
	return
}