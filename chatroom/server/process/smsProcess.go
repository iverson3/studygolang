package process

import (
	"encoding/json"
	"fmt"
	"net"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/utils"
)

type SmsProcess struct {
	Conn net.Conn
}

func (this *SmsProcess) SendGroupSmsMessage(mes *common.Message) (err error) {
	var groupSms common.GroupSmsMes
	err = json.Unmarshal([]byte(mes.Data), &groupSms)
	if err != nil {
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		return
	}

	MyUserManager.rwMutex.RLock()
	for uid, up := range MyUserManager.onlineUsers {
		if uid == groupSms.UserId {
			continue
		}
		err2 := this.SendSmsToOnlineUser(data, up.Conn)
		if err2 != nil {
			fmt.Println("发送失败，userid: ", uid)
		}
	}
	MyUserManager.rwMutex.RUnlock()
	return
}

// 发送聊天消息给指定的用户 (群聊 私聊通用)
func (this *SmsProcess) SendSmsToOnlineUser(data []byte, conn net.Conn) (err error) {
	tf := &utils.Transfer{Conn: conn}
	return tf.WritePkg(data)
}

func (this *SmsProcess) SendPersonalSmsMessage(mes *common.Message) (err error) {
	var personalSms common.PersonalSmsMes
	err = json.Unmarshal([]byte(mes.Data), &personalSms)
	if err != nil {
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		return
	}

	up, err := MyUserManager.GetOnlineUserById(personalSms.To.UserId)
	if err != nil {
		return
	}

	err = this.SendSmsToOnlineUser(data, up.Conn)
	return
}

























