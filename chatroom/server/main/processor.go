package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
	processes "studygolang/chatroom/server/process"
	"studygolang/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) mainProcess() (err error) {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		fmt.Println("等待客户端发送数据~~~")
		message, err := tf.ReadPkg()
		if err != nil {
			// 判断错误是否是客户端关闭了连接
			if err == io.EOF || strings.Contains(err.Error(), "close") {
				err = model.ERROR_CLIENT_CLOSE_CONNECTION
			}

			// 客户端断开连接，则将当前用户从在线用户列表中移除
			//_, _ = process2.RemoveUserFromOnlineUserList()
			return err
		}

		err = this.ServerProcessMess(&message)
	}
}

// 根据客户端消息的类型，调用不同的处理函数进行业务处理
func (this *Processor) ServerProcessMess(mess *common.Message) (err error) {
	switch mess.Type {
	case common.LoginMesType:
		// 登录处理
		up := &processes.UserProcess{Conn: this.Conn}
		return up.ServerProcessLogin(mess)
	case common.RegisterMesType:
		// 注册处理
		up := &processes.UserProcess{Conn: this.Conn}
		return up.ServerProcessRegister(mess)
	default:
		fmt.Println("message: ", mess)
		return errors.New("未知的消息类型：" + mess.Type)
	}
}
