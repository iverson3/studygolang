package process

import (
	"encoding/json"
	"fmt"
	common "studygolang/chatroom/common/message"
)

func ShowGroupSms(mess *common.Message) error {
	var groupSms common.GroupSmsMes
	err := json.Unmarshal([]byte(mess.Data), &groupSms)
	if err != nil {
		return err
	}

	fmt.Println("===========================")
	fmt.Printf("用户[%s]群发了消息：%s (%s)\n", groupSms.UserName, groupSms.Content, groupSms.SendTime[:19])
	fmt.Println("===========================")
	return nil
}

func ShowPersonalSms(mess *common.Message) error {
	var personalSms common.PersonalSmsMes
	err := json.Unmarshal([]byte(mess.Data), &personalSms)
	if err != nil {
		return err
	}

	fmt.Println("===========================")
	fmt.Printf("用户[%s]给你私发了消息：%s (%s)\n", personalSms.From.UserName, personalSms.Content, personalSms.SendTime[:19])
	fmt.Println("===========================")
	return nil
}


