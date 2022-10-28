package main

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to server success")
	handleMsg(conn)
}

func handleMsg(conn net.Conn) {
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	exitCh := make(chan struct{}, 1)
	stopCh := make(chan struct{}, 1)
	go closeConn(conn, exitCh, stopCh)

	var err error
	var failCount int
	for {
		select {
		case <-stopCh:
			return
		default:
			// 连续错误超过三次则认为连接不再可用
			if failCount > 3 {
				exitCh <- struct{}{}
				fmt.Println("fail too many times, break for")
				break
			}

			message := "i'm client msg."
			err = wsutil.WriteClientMessage(conn, ws.OpText, []byte(message))
			if err != nil {
				failCount++
				fmt.Errorf("wsutil.WriteClientMessage() failed, error: %v\n", err)
				panic(err)
			}
			fmt.Println("send message to server success")

			data, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				failCount++
				fmt.Errorf("wsutil.ReadServerData() failed, error: %v\n", err)
				continue
			}

			failCount = 0
			fmt.Println("got msg from server: ", string(data))
			time.Sleep(5 * time.Second)
		}
	}
}

func closeConn(conn net.Conn, exitCh chan struct{}, stopCh chan struct{}) {
	ch := make(chan os.Signal)
	// 监听指定的信号
	signal.Notify(ch, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	// 阻塞直到有信号到来 或者 退出通知的到来
	select {
	case <-ch:
		stopCh <- struct{}{}
		// 释放连接资源
		if conn != nil {
			// 程序退出之前给服务器发送一个关闭连接的消息
			message := "#close-connection#"
			err := wsutil.WriteClientMessage(conn, ws.OpText, []byte(message))
			if err != nil {
				fmt.Errorf("close-msg send failed, error: %v", err)
			}
			conn.Close()
		}

		fmt.Println("client exit.")
	case <-exitCh:
		// 停止监听信号
		signal.Stop(ch)
		return
	}
}