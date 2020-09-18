package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"studygolang/netProxyServer/utils"
)

// 实现反向代理服务器 - 负载均衡算法

type BalanceProxyHandler struct {}

func (*BalanceProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			writer.WriteHeader(500)
			fmt.Println("server error: ", err)
		}
	}()
	// chrome会访问一个图标文件，忽略
	if request.URL.Path == "/favicon.ico" {
		return
	}

	// 利用随机算法选出server
	//server := utils.LB.SelectServerByRand()

	// 利用ip hash算法选出server
	//server := utils.LB.SelectServerByIPHash(request.RemoteAddr)

	// 利用加权随机算法选出server
	//server := utils.LB.SelectServerByWeightRand()

	// 利用改良版的加权随机算法选出server
	//server := utils.LB.SelectServerByWeightRand2()

	// 利用简单轮询算法选出server
	//server := utils.LB.SelectServerByRoundRobin()

	// 利用平滑加权轮询算法选出server
	//server := utils.LB.SelectServerBySmoothWeightRoundRobin()

	// 平滑加权轮询算法实现负载均衡 + 降权机制
	server := utils.LB.SelectServerBySmoothWeightRoundRobinWithFailWeight()


	target, _ := url.Parse(server.Host)
	// 使用go内置的反向代理函数实现反向代理功能
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(writer, request)
}

func main() {
	// 构建反向代理服务器
	_ = http.ListenAndServe(":8089", &BalanceProxyHandler{})
}
