package process

import (
	"encoding/json"
	"studygolang/chatroom/client/utils"
	common "studygolang/chatroom/common/message"
	"time"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupSms(content string) (err error) {
	var mess common.Message
	mess.Type = common.GroupSmsMesType

	// 组装数据
	var groupSms common.GroupSmsMes
	groupSms.UserId     = CurUser.UserId
	groupSms.UserName   = CurUser.UserName
	groupSms.UserStatus = CurUser.UserStatus
	groupSms.Content    = content
	groupSms.SendTime   = time.Now().String()

	bytes, err := json.Marshal(groupSms)
	if err != nil {
		return
	}

	mess.Data = string(bytes)
	data, err := json.Marshal(mess)
	if err != nil {
		return
	}

	tf := &utils.Transfer{Conn: CurUser.Conn}
	return tf.WritePkg(data)
}