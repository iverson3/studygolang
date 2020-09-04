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

				// 客户端断开连接，则先通知其他用户,当前用户下线的消息
				up, _ := processes.MyUserManager.GetOnlineUserByAddr(this.Conn.RemoteAddr().String())
				if up == nil {
					return err
				}
				user := &model.User{
					UserId:     up.UserId,
					UserName:   up.UserName,
				}
				up.NotifyOtherOfflineUsers(user)

				// 客户端断开连接，则将当前用户从在线用户列表中移除
				processes.MyUserManager.DelOnlineUserByAddr(this.Conn.RemoteAddr().String())
				return err
			}
			continue
		}

		// 可以考虑使用go携程来处理
		// 或者开一个携程先处理并组装数据 再向消息channel里写入需要发送的数据，然后由另一个运行中的携程来从channel中取出数据 发给指定的用户(群发或私发)
		err = this.ServerProcessMess(&message)
		if err != nil {
			fmt.Printf("处理客户端消息出错！error: %s \n", err.Error())
		}
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
	case common.GroupSmsMesType:
		up := &processes.SmsProcess{Conn: this.Conn}
		return up.SendGroupSmsMessage(mess)
	case common.PersonalSmsMesType:
		up := &processes.SmsProcess{Conn: this.Conn}
		return up.SendPersonalSmsMessage(mess)
	case common.FileUploadMesType:
		up := &processes.FileProcess{}
		return up.UploadBigFile(mess)
	default:
		fmt.Println("message: ", mess)
		return errors.New("未知的消息类型：" + mess.Type)
	}
}
