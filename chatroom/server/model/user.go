package model

type User struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
	UserStatus string `json:"user_status"`   // 用户状态 (在线 离线)
}

