package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 测试是否能正常引用到静态资源文件
	test, err := os.Open("/usr/local/resources/img/test.png")
	if err != nil {
		fmt.Println("open file failed")
		fmt.Println(err)
		return
	}
	stat, err := test.Stat()
	if err == nil {
		fmt.Println(stat.Name())
		fmt.Println(stat.Size())
	}
	test.Close()


	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("start to accept connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("accept a connection")
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	fmt.Println("waiting msg from client...")
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Printf("got msg from cient: %s", msg)

		resp := fmt.Sprintf("%s from server\n", msg[:len(msg)-1])
		_, _ = conn.Write([]byte(resp))
	}
}

