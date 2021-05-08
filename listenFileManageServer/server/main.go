package main

import (
	"encoding/json"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
)

type Conf struct {
	Port int `json:"port"`
}

const (
	confFilePath = "../config/config.json"
)

func main() {
	data, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Error("read config file failed! error: ", err)
		return
	}

	var conf Conf
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Error("unmarshal config data failed! error: ", err)
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Error("listen failed! error: ", err)
		return
	}

	fmt.Printf("server listen in port:%d...\n", conf.Port)

	go listenExitSignal()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("got a connect from client")

		go func(conn net.Conn) {
			defer conn.Close()

			_, _ = conn.Write([]byte("hello\n"))
		}(conn)
	}
}

func listenExitSignal()  {
	ch := make(chan os.Signal)

	// 获取程序退出信号
	signal.Notify(ch, os.Interrupt, os.Kill)
	signals := <-ch

	fmt.Printf("got signal: %v\n", signals)

	fmt.Println("server exit!")
	os.Exit(1)
}

