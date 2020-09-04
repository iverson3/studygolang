package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func RequestUrl(writer http.ResponseWriter, request *http.Request, url string) (err error) {
	// 构建一个新的http请求 (请求代理服务器所代理的真实的服务器)
	newRequest, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		return
	}
	// 将代理服务器request里面的header信息拷贝到新的http请求header里面去
	CloneHeader(request.Header, &newRequest.Header)
	// 增加新的header头信息，以便真实服务器中能凭此获取到真实的客户端IP地址
	newRequest.Header.Add("x-forwarded-for", request.RemoteAddr)

	// 用http包中自定义Transport去执行构建的请求(请求真实服务器) 并获得请求的响应结果
	dt := &http.Transport{
		DialContext:            (&net.Dialer{
			Timeout:       30 * time.Second,
			KeepAlive:     30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout:  1 * time.Second,
	}
	response, err := dt.RoundTrip(newRequest)

	// 用http包中默认的Transport去执行构建的请求(请求真实服务器) 并获得请求的响应结果
	// response, err := http.DefaultTransport.RoundTrip(newRequest)

	// 用http包中默认的Client去执行构建的请求(请求真实服务器) 并获得请求的响应结果
	// response, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		return
	}
	defer response.Body.Close()

	writerHeader := writer.Header()
	// 将响应中的header信息拷贝到代理服务器的writer中，以便一同发送给客户端
	CloneHeader(response.Header, &writerHeader)
	// 将响应中的http状态码写入到代理服务器的writer中，以便一同发送给客户端
	writer.WriteHeader(response.StatusCode)

	// 读取响应结果中的所有内容
	bytes, err := ioutil.ReadAll(response.Body)
	// 将响应数据返回给客户端(浏览器)
	_, err = writer.Write(bytes)
	return
}

// 拷贝http Header头信息
func CloneHeader(src http.Header, dest *http.Header) {
	for k, v := range src {
		dest.Set(k, v[0])
	}
}

// 获取客户端的真实IP地址
// request.RemoteAddr可能是代理服务器的ip地址 (如果有代理服务器作为中间代理)
func GetRealClientIp(request *http.Request) string {
	ips := request.Header.Get("x-forwarded-for")
	if ips != "" {
		ipSlice := strings.Split(ips, ",")
		if len(ipSlice) > 0 && ipSlice[0] != "" {
			return ipSlice[0]
		}
	}
	return request.RemoteAddr
}
