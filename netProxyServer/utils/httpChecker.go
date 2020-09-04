package utils

import (
	"math"
	"net/http"
	"time"
)

// 计数器机制 vs 降权机制

type HttpChecker struct {
	Servers HttpServers
	FailMax int    // 服务器检测时，服务器连续响应失败的最大次数 (达到该次数之后 服务器就会被设置为DOWN)
	RecoverMax int // 服务器检测时，服务器连续响应成功的最大次数 (达到该次数之后 服务器就会被设置为UP)
	FailFactor float64 // 降权因子
}

func NewHttpChecker(servers HttpServers, failMax int, recoverMax int, failFactor float64) *HttpChecker {
	return &HttpChecker{Servers: servers, FailMax: failMax, RecoverMax: recoverMax, FailFactor: failFactor}
}

// 监测服务器是否可用
// 一般做法是发送Head请求 (HEAD/GET/POST/PUT/DELETE)
// 使用Head请求，服务器只返回头部信息，不会返回body部分，网络传输量比较小，一般用来检测url是否可用
func (this *HttpChecker) Check(timeout time.Duration) {
	client := http.Client{}
	for _, server := range this.Servers {
		resp, err := client.Head(server.Host)
		if resp != nil {
			defer resp.Body.Close()
		}

		if err != nil {
			this.responseFail(server)
			continue
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			this.responseSuccess(server)
		} else {
			this.responseFail(server)
		}
	}
}

func (this *HttpChecker) responseFail(server *HttpServer) {
	// 计数器机制使用以下代码
	// 服务器响应失败的次数达到了阈值 则判定服务器为不可用状态
	if server.FailCount >= this.FailMax {
		server.Status = "DOWN"
	} else {
		server.FailCount++
	}
	server.SuccessCount = 0


	// 降权机制使用以下代码
	fw := int(math.Floor(float64(server.Weight) * (1 / this.FailFactor)))
	if fw == 0 {
		fw = 1
	}
	// 服务器每响应失败一次，降权值就加一次 原始权重*(1/降权因子)
	server.FailWeight += fw
	// FailWeight不能超过Weight
	if server.FailWeight > server.Weight {
		server.FailWeight = server.Weight
	}
}
func (this *HttpChecker) responseSuccess(server *HttpServer) {
	// 计数器机制使用以下代码
	if server.FailCount > 0 {
		server.FailCount--
		server.SuccessCount++

		// 服务器响应成功的次数达到了阈值 则判定服务器为可用状态
		if server.SuccessCount >= this.RecoverMax {
			server.Status = "UP"
			server.SuccessCount = 0
			server.FailCount = 0
		}
	} else {
		server.Status = "UP"
	}


	// 降权机制使用以下代码
	// 一旦服务器响应成功一次，则直接将降权值还原为初始值 0
	server.FailWeight = 0
}