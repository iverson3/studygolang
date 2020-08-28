package main

import (
	"fmt"
	"os"
	"studygolang/chatroom/client/process"
)

// 用户账号和密码
var userId int
var userName string
var userPwd string

func main() {
	var key int

	fmt.Println("---------------欢迎来到多人聊天系统-----------------")
	fmt.Println("\t\t 1.登录聊天室")
	fmt.Println("\t\t 2.注册用户")
	fmt.Println("\t\t 3.退出系统")

	for {
		fmt.Println("\t\t 请选择功能 (1-3)：")
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入用户ID：")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			_, _ = fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("login failed! err: ", err)
			} else {
				fmt.Println("login success")
			}
		case 2:
			fmt.Println("请输入用户ID：")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称：")
			_, _ = fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			err := up.Register(userId, userName, userPwd)
			if err != nil {
				fmt.Println("register failed! err: ", err)
			} else {
				fmt.Println("register success")
			}
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("选择有误，请重新进行选择")
		}
	}
}
