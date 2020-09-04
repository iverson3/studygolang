package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"studygolang/netProxyServer/utils"
)

type web1Handler struct {}

// 实现Handler接口
func (web1Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Authorization：Basic cAkcSKDqwJGJff===
	auth := request.Header.Get("Authorization")
	if auth == "" {
		// 通过在header中设置WWW-Authenticate,实现Basic Auth认证
		writer.Header().Set("WWW-Authenticate", `Basic realm="您必须输入用户名和密码"`)
		// 返回401状态码，表示未通过认证的请求
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	// Basic cAkcSKDqwJGJff===
	authSlice := strings.Split(auth, " ")
	if len(authSlice) == 2 && authSlice[0] == "Basic" {
		// base64解码得到客户端用户填写的用户名和密码 (格式："username:password")
		bytes, err := base64.StdEncoding.DecodeString(authSlice[1])
		if err == nil && string(bytes) == "admin:123456" {
			writer.Write([]byte(fmt.Sprintf("<h1>web1,from: %s</h1>", utils.GetRealClientIp(request))))
			return
		}
	}

	writer.Write([]byte("用户名或密码有误"))
}

type web2Handler struct {}

// 实现Handler接口
func (web2Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("web2"))
}

type web3Handler struct {}

// 实现Handler接口
func (web3Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("web3"))
}

func main() {
	// 操作系统信号channel
	c := make(chan os.Signal)

	go func() {
		err := http.ListenAndServe(":8881", web1Handler{})
		if err != nil {
			fmt.Println("linsten failed! error: ", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":8882", web2Handler{})
		if err != nil {
			fmt.Println("linsten failed! error: ", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":8883", web3Handler{})
		if err != nil {
			fmt.Println("linsten failed! error: ", err)
		}
	}()

	// Notify函数让signal包将输入信号转发到channel
	signal.Notify(c, os.Interrupt)
	s := <-c

	fmt.Println("s: ", s)
}
