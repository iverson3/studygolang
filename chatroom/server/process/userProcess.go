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
		userManager.AddOnlineUser(this)
		go this.NotifyOtherOnlineUsers(curUser)

		// 构建在线用户列表数据，返回给客户端
		userList := make(map[int]string)
		users := userManager.GetAllOnlineUser()

		for _, user := range users {
			userList[user.UserId] = user.UserName
		}
		loginResMess.OnlineUserList = userList
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
	data, err := assembleNotifyUserOnlineData(user)
	if err != nil {
		return
	}

	// 将用户上线的通知信息数据广播给所有已在线的用户
	for uid, up := range userManager.onlineUsers {
		if uid != user.UserId {
			err = up.NotifyMeOnline(data)
			if err != nil {
				fmt.Println("通知用户上线失败，失败的用户: ", uid)
			}
		}
	}
}

func assembleNotifyUserOnlineData(user *model.User) (data []byte, err error) {
	var mess common.Message
	mess.Type = common.NotifyUserStatusMesType

	var notifyMess common.NotifyUserStatusMes
	notifyMess.UserId = user.UserId
	notifyMess.UserName = user.UserName
	notifyMess.Status = common.UserOnline

	bytes, err := json.Marshal(notifyMess)
	if err != nil {
		return
	}
	mess.Data = string(bytes)

	return json.Marshal(mess)
}

func (this *UserProcess) NotifyMeOnline(data []byte) (err error) {
	tf := &utils.Transfer{Conn: this.Conn}
	return tf.WritePkg(data)
}