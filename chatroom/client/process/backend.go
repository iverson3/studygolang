package process

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"studygolang/chatroom/client/utils"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
)

func ShowMenu() {
	fmt.Println("---------------聊天系统主界面----------------")
	fmt.Println("\t\t 1.显示在线用户列表")
	fmt.Println("\t\t 2.群发消息")
	fmt.Println("\t\t 3.个人私聊")

	fmt.Println("\t\t 4.文件上传")
	fmt.Println("\t\t 5.信息列表")
	fmt.Println("\t\t 6.退出登录")

	fmt.Println("\t\t 7.退出系统")

	smsProcess := &SmsProcess{}
	var key int
	for {
		fmt.Println("\t\t 请选择 (1-7)：")
		// 需要处理用户输入字符串的情况
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			ShowOnlineUserList()
		case 2:
			var content string
			fmt.Println("请输入需要群发的消息内容:")
			_, _ = fmt.Scanf("%s\n", &content)

			err := smsProcess.SendGroupSms(content)
			if err != nil {
				fmt.Println("群发消息失败；error: ", err.Error())
			}
		case 3:
			var uid int
			var content string
			fmt.Println("请输入私聊用户的Id:")
			for {
				_, _ = fmt.Scanf("%d\n", &uid)
				if ExistUserInUserList(uid) && uid != CurUser.UserId {
					break
				}
				if !ExistUserInUserList(uid) {
					fmt.Println("该用户不存在或不在线，请重新输入私聊用户的Id:")
				}
				if uid == CurUser.UserId {
					fmt.Println("不能跟自己私聊，请重新输入私聊用户的Id:")
				}
			}
			fmt.Println("请输入需要发送的消息内容:")
			_, _ = fmt.Scanf("%s\n", &content)

			err := smsProcess.SendPersonalSms(uid, content)
			if err != nil {
				fmt.Println("私发消息失败；error: ", err.Error())
			}
		case 4:
			var path string
			fmt.Print("请输入文件路径:")
			_, _ = fmt.Scanf("%s\n", &path)

			info, err := os.Stat(path)
			if err == nil || os.IsExist(err) {
				if info.IsDir() {
					fmt.Println("不是文件 是目录")
				} else {
					fp := FileProcess{}
					err = fp.UploadBigFile(path)
				}
			} else {
				fmt.Println("文件不存在")
			}
		case 5:
			fmt.Println("5")
		case 6:
			fmt.Println("6")
		case 7:
			fmt.Println("7")
			os.Exit(0)
		default:
			println("输入有误，请重新输入")
		}
	}
}

// 在后台保持跟服务端之间的通讯
// 随时准备接收服务端发送过来的数据 (比如用户上线下线提醒 消息推送提醒等等)
func processServerMess(conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}
	for {
		mess, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "close") {
				err = model.ERROR_SERVER_CLOSE_CONNECTION
			}
			fmt.Println("接收服务端的消息出错，error：", err.Error())

			// 尝试重连服务器，恢复与服务器的通讯
			// 服务器断开连接 有两种情况： 服务器停止运行 | 网络故障意外断开了连接
			return
		}

		// 根据消息内容 进行不同的逻辑处理
		switch mess.Type {
		case common.NotifyUserStatusMesType:
			fmt.Println("收到来自服务端的好友上下线通知消息")
			var notifyMess common.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mess.Data), &notifyMess)
			if err != nil {
				fmt.Println("反序列化消息出错")
			} else {
				UpdateOnlineUserList(&notifyMess)
			}
		case common.GroupSmsMesType:
			err := ShowGroupSms(&mess)
			if err != nil {
				fmt.Println("显示群发消息失败！error: ", err.Error())
			}
		case common.PersonalSmsMesType:
			err := ShowPersonalSms(&mess)
			if err != nil {
				fmt.Println("显示私发消息失败！error: ", err.Error())
			}
		default:
			fmt.Println("未知的消息类型")
		}
	}
}