package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const url = "http://www.imooc.com"

// 使用http client发送请求
func main() {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	// 设置请求header信息
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	// 自定义client ( 默认 http.DefaultClient.Do() )
	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			// 返回nil 表示允许重定向
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
	resp, err := client.Do(request)

	//resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))
}
