package process

import (
	"encoding/json"
	"errors"
	"studygolang/chatroom/client/utils"
	common "studygolang/chatroom/common/message"
	"studygolang/chatroom/server/model"
	"time"
)

type SmsProcess struct {

}

// 发送群消息
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

	return jsonAndSendData(mess, groupSms)
}

// 发送私聊消息
func (this *SmsProcess) SendPersonalSms(userId int, content string) (err error) {
	var mess common.Message
	mess.Type = common.PersonalSmsMesType

	// 组装数据
	fromUser := model.User{
		UserId:     CurUser.UserId,
		UserName:   CurUser.UserName,
	}
	user := GetUserFromOnlineUserList(userId)
	if user == nil {
		return errors.New("私聊的用户不存在或已下线")
	}
	toUser := model.User{
		UserId:     userId,
		UserName:   user.UserName,
	}

	var personalSms common.PersonalSmsMes
	personalSms.From     = fromUser
	personalSms.To       = toUser
	personalSms.Content  = content
	personalSms.SendTime = time.Now().String()

	return jsonAndSendData(mess, personalSms)
}

func jsonAndSendData(mess common.Message, sms interface{}) (err error) {
	bytes, err := json.Marshal(sms)
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