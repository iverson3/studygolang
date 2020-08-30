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

	userManager.rwMutex.RLock()
	for uid, up := range userManager.onlineUsers {
		if uid == groupSms.UserId {
			continue
		}
		err2 := this.SendGroupSmsToEachOnlineUser(data, up.Conn)
		if err2 != nil {
			fmt.Println("发送失败，userid: ", uid)
		}
	}
	userManager.rwMutex.RUnlock()
	return
}

func (this *SmsProcess) SendGroupSmsToEachOnlineUser(data []byte, conn net.Conn) (err error) {
	tf := &utils.Transfer{Conn: conn}
	return tf.WritePkg(data)
}


























