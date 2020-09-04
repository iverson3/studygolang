package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"studygolang/netProxyServer/utils"
)

// 实现反向代理服务器

type ProxyHandler struct {}

func (*ProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)

	defer func() {
		if err := recover(); err != nil {
			writer.WriteHeader(500)
			fmt.Println("server error: ", err)
		}
	}()

	// 根据匹配到的不同path，将请求分发给对应的真实服务器
	// 实现代理服务器的请求转发
	for path, passInfo := range utils.ProxyConfigs {
		matched, _ := regexp.MatchString(path, request.URL.Path)
		if matched {
			// 实现反向代理请求真实服务器
			//_ = utils.RequestUrl(writer, request, passInfo.Url)

			// 使用go内置的反向代理函数实现反向代理功能
			target, _ := url.Parse(passInfo.Url)
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(writer, request)
			return
		}
	}
	// _, _ = writer.Write([]byte("404"))

	_, _ = writer.Write([]byte("proxy server: default index page"))
}

func main() {
	// 构建反向代理服务器
	_ = http.ListenAndServe(":8089", &ProxyHandler{})
}
