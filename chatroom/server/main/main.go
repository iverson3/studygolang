package main

import (
	"fmt"
	"net"
	"studygolang/chatroom/server/model"
	"time"
)

// 处理与客户端之间的通讯
func process(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{Conn: conn}
	err := processor.mainProcess()
	if err != nil {
		fmt.Println("服务器与客户端通讯出现错误，error: ", err.Error())
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen failed! error: ", err)
		return
	}
	defer listener.Close()

	// 服务端的监听启动之后，初始化redis连接池
	initPool("127.0.0.1:6379", "13396095889", 16, 0, 300 * time.Second)
	defer pool.Close()

	// 实例化全局的UserDao
	initUserDao()

	for {
		fmt.Println("等待客户端的连接~~~")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept failed! error: ", err)
			continue
		}

		go process(conn)
	}
}
