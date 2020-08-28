package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("tcp connect to server failed! error: ", err)
		return
	}
	defer conn.Close()

	// 获取键盘输入 并发送给server
	reader := bufio.NewReader(os.Stdin)  // os.Stdin 标准输入终端

	for {
		fmt.Println("请输入内容发送到服务器：(exit表示退出)")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取终端输入失败！ error: ", err)
			continue
		}

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出!")
			break
		}
		if line == "" {
			fmt.Println("发送的内容不能为空!")
			continue
		}

		line = line + "\n"
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("客户端发送数据失败！ error: ", err)
			continue
		}
		fmt.Println("客户端发送成功！ 内容: ", line)
	}
}
