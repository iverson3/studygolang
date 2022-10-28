package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"io/ioutil"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	upgrader := ws.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Protocol:        nil,
		ProtocolCustom:  nil,
		ExtensionCustom: nil,
		Negotiate:       nil,
		Header:          nil,
		OnRequest:       nil,
		OnHost:          nil,
		OnHeader:        nil,
		OnBeforeUpgrade: nil,
	}

	fmt.Println("start to accept connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("listener.Accept() failed, error: %v\n", err)
			continue
		}

		fmt.Println("got a new connection")

		// 对连接进行升级，升级为websocket
		if _, err := upgrader.Upgrade(conn); err != nil {
			fmt.Errorf("upgrader.Upgrade() failed, error: %v\n", err)
			continue
		}

		fmt.Println("connection upgrade success")

		go handleMsg(conn)
	}
}

//func heartbeat(conn net.Conn, closeCh chan struct{}) {
//	tick := time.Tick(3 * time.Second)
//	var failTimes int
//
//	for {
//		// 连续三次心跳失败则认为客户端连接已关闭
//		if failTimes >= 3 {
//			closeCh <- struct{}{}
//			return
//		}
//		select {
//		case <-tick:
//			res := "ping"
//			err := wsutil.WriteServerText(conn, []byte(res))
//			if err != nil {
//				failTimes++
//				fmt.Errorf("wsutil.WriteServerText() failed, error: %v\n", err)
//				continue
//			}
//
//			reader := wsutil.NewReader(conn, ws.StateServerSide)
//			_, err = reader.NextFrame()
//			if err != nil {
//				failTimes++
//				fmt.Errorf("reader.NextFrame() failed, error: %v\n", err)
//				continue
//			}
//
//			data, err := ioutil.ReadAll(reader)
//			if err != nil {
//				failTimes++
//				fmt.Errorf("ioutil.ReadAll() from reader failed, error: %v\n", err)
//				continue
//			}
//
//			fmt.Println("[heartbeat] got msg from client: ", string(data))
//
//			if string(data) != "pong" {
//				failTimes++
//			} else {
//				failTimes = 0
//			}
//		}
//	}
//}

func handleMsg(conn net.Conn) {
	defer conn.Close()

	closeConnCh := make(chan struct{}, 1)
	//go heartbeat(conn, closeConnCh)

	fmt.Println("start to handle msg")
	// todo: 客户端连接关闭了，如何感知，方便服务端做善后处理
	// todo: 服务器关闭或宕机了，做好善后处理
	for {

		//header, err := ws.ReadHeader(conn)
		//if header.OpCode == ws.OpClose {
		//	fmt.Println("client closed the connection")
		//	return
		//}

		select {
		case <-closeConnCh:
			fmt.Println("client connection had closed")
			return
		default:
			reader := wsutil.NewReader(conn, ws.StateServerSide)
			_, err := reader.NextFrame()
			if err != nil {
				fmt.Errorf("reader.NextFrame() failed, error: %v\n", err)
				continue
			}

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				fmt.Errorf("ioutil.ReadAll() from reader failed, error: %v\n", err)
				continue
			}

			// 判断是否是客户端关闭连接的消息
			if string(data) == "#close-connection#" {
				fmt.Println("client connection has closed")
				return
			} else {
				fmt.Println("got msg from client: ", string(data))
			}

			res := "i'm a msg from server."
			err = wsutil.WriteServerText(conn, []byte(res))
			if err != nil {
				fmt.Errorf("wsutil.WriteServerText() failed, error: %v\n", err)
				continue
			}

			fmt.Println("send msg to client success")
		}
	}
}