package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("accept new connection")
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

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

