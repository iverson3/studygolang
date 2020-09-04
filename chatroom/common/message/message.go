package common

import "studygolang/chatroom/server/model"

// 消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	GroupSmsMesType         = "GroupSmsMes"
	PersonalSmsMesType      = "PersonalSmsMes"
	FileUploadMesType       = "FileUploadMes"
)

const (
	UserOnline = "在线"
	UserOffline = "离线"
	UserBusy = "忙碌"
)

type Message struct {
	Type string `json:"type"`  // 消息类型
	Data string `json:"data"`  // 消息内容
	AddTime string `sql:"add_time"`  // 演示 sql tag的用法
}
type ResponseMes struct {
	Code int     `json:"code"`   // 状态码
	Error string `json:"error"`  // 错误信息
}

// 登录消息
type LoginMes struct {
	UserId int      `json:"user_id"`
	UserPwd string  `json:"user_pwd"`
	UserName string `json:"user_name"`
}
// 登录结果消息
type LoginResMes struct {
	ResponseMes
	UserName string `json:"user_name"`
	OnlineUserList map[int]string  // 在线用户列表数据 ( 结构: map[userId]userName )
}

// 注册消息
type RegisterMes struct {
	//User model.User `json:"user"`
	model.User
}
// 注册结果消息
type RegisterResMes struct {
	ResponseMes
}

// 服务端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	Status string `json:"status"`  // 用户状态
}

// 群聊天消息
type GroupSmsMes struct {
	model.User
	Content string `json:"content"`
	SendTime string `json:"send_time"`
}
// 个人聊天消息
type PersonalSmsMes struct {
	From model.User `json:"from"`
	To model.User `json:"to"`
	Content string `json:"content"`
	SendTime string `json:"send_time"`
}

// 文件上传消息
type FileUploadMes struct {
	FileName string `json:"file_name"`
	Data []byte `json:"data"`
	Len int `json:"len"`
	UploadStatus string `json:"upload_status"`
}

// 获取在线用户列表结果消息
//type GetOnlineUserListResMes struct {
//	ResponseMes
//	UserList []model.User `json:"user_list"`
//}
