package main

import (
	"fmt"
	"net"
	"strings"
)

func Worker(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		fmt.Println("等待客户端发送数据过来~~~")
		// 等待并读取客户端发送过来的数据
		n, err := conn.Read(data)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				fmt.Println("客户端已断开连接!")
				return
			} else {
				fmt.Println("read data failed! error: ", err)
			}
		} else {
			fmt.Printf("客户端[%s]发送的数据： %s", conn.RemoteAddr().String(), string(data[:n]))
		}
	}
}

func main() {
	fmt.Println("服务器在8889端口开始监听~~~")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("tcp listen failed! error: ", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端的连接~~~")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("tcp listen accept failed! error: ", err)
		} else {
			fmt.Println("客户端连接成功~~~")
			fmt.Println("client ip address: ", conn.RemoteAddr().String())
			go Worker(conn)
		}
	}

}
