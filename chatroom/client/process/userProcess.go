package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"studygolang/chatroom/client/utils"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
)

var (
	serverAddress = "127.0.0.1:8889"
)

type UserProcess struct {
	
}

func connectServer() (conn net.Conn, err error) {
	// 跟服务器建立连接并通讯
	return net.Dial("tcp", serverAddress)
}

// 实现客户端与服务端的登录请求交互
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := connectServer()
	if err != nil {
		return
	}
	defer conn.Close()

	// 构建消息结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	bytes, err := json.Marshal(loginMes)
	if err != nil {
		return
	}

	var mess common.Message
	mess.Type = common.LoginMesType
	mess.Data = string(bytes)

	data, err := json.Marshal(mess)
	if err != nil {
		return
	}

	// 向服务端发送用户登录数据
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	// 接收服务器发送过来的数据
	message, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 反序列化服务器发过来的数据，得到登录的结果信息
	var loginRes common.LoginResMes
	err = json.Unmarshal([]byte(message.Data), &loginRes)
	if err != nil {
		return
	}

	// 根据Code码判断登录的结果
	if loginRes.Code == 200 {
		fmt.Println("login success!")

		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserName = loginRes.UserName
		CurUser.UserStatus = common.UserOnline

		for i, v := range loginRes.OnlineUserList {
			// 将服务端返回的在线用户列表放入本地维护的好友列表中
			user := &model.User{
				UserId:     i,
				UserName:   v,
				UserStatus: common.UserOnline,
			}
			onlineUserList[i] = user
		}

		go processServerMess(conn)
		ShowMenu()
		return
	} else {
		return errors.New(loginRes.Error)
	}
}

func (this *UserProcess) Register(userId int, userName, userPwd string) (err error) {
	conn, err := connectServer()
	if err != nil {
		return
	}
	defer conn.Close()

	// 构建消息结构体
	var regMes common.RegisterMes
	regMes.UserId = userId
	regMes.UserName = userName
	regMes.UserPwd = userPwd
	
	bytes, err := json.Marshal(regMes)
	if err != nil {
		return
	}

	var mess common.Message
	mess.Type = common.RegisterMesType
	mess.Data = string(bytes)

	data, err := json.Marshal(mess)
	if err != nil {
		return
	}

	// 向服务端发送用户注册数据
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	// 接收服务器发送过来的数据
	message, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 反序列化服务器发过来的数据，得到注册的结果信息
	var loginRes common.RegisterResMes
	err = json.Unmarshal([]byte(message.Data), &loginRes)
	if err != nil {
		return
	}

	// 根据Code码判断登录的结果
	if loginRes.Code == 200 {
		return
	} else {
		return errors.New(loginRes.Error)
	}
}