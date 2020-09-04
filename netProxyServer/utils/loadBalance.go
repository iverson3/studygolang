package utils

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"sort"
	"time"
)

var LB *LoadBalance

// 真实服务器结构体
type HttpServer struct {
	Host string
	Weight int        // 权重
	CurWeight int     // 当前权重 (平滑加权轮询算法需要使用)
	FailWeight int    // 降权值，初始值为0 (降权机制下使用; 服务器每响应失败一次，该降权值就增加一个数[比如 原始权重*(1/降权因子)])
	Status string     // 服务器状态，默认为UP ("UP":可用 "DOWN":宕机)
	FailCount int     // 计数器 (在服务器检测中，服务器连续响应失败的次数)
	SuccessCount int  // 计数器 (在服务器检测中，服务器连续响应成功的次数)
}
type HttpServers []*HttpServer
// 下面实现Sort接口的三个必要方法，以便能对LoadBalance里面的servers(HttpServers类型)进行排序
func (hs HttpServers) Len() int {
	return len(hs)
}
func (hs HttpServers) Less(i, j int) bool {
	// 排序的依据是按照CurWeight 当前权重进行排序
	return hs[i].CurWeight > hs[j].CurWeight
}
func (hs HttpServers) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight, CurWeight: weight, Status: "UP"}
}

// 负载均衡
type LoadBalance struct {
	Servers HttpServers
	CurIndex int  // 当前访问的服务器下标
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HttpServer, 0, )}
}

var ServerIndexes []int  // 结构类似 [0, 0, 0, 1, 1, 1, 1, 1]
var totalWeight int      // 所有服务器的总权重

// 初始化
func init() {
	LB = NewLoadBalance()
	// 手动添加两个真实服务器
	//LB.AddServer(NewHttpServer("http://127.0.0.1:8881"))
	//LB.AddServer(NewHttpServer("http://127.0.0.1:8882"))

	// 通过读取配置文件 获取所有可用的服务器地址
	for _, passInfo := range ProxyConfigs {
		LB.AddServer(NewHttpServer(passInfo.Url, passInfo.Weight))
	}

	ServerIndexes = make([]int, 1)
	// 根据服务器权重的不同 设置不同数目的"抽奖球" 以达到权重的效果
	for index, server := range LB.Servers {
		totalWeight += server.Weight
		// 权重设置为0的服务器不参与负载均衡
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndexes = append(ServerIndexes, index)
			}
		}
	}

	// 从配置文件中读取
	failMax := 3        // 服务器无响应次数阈值 (连续无响应的次数)
	recoverMax := 2     // 服务器可用状态恢复阈值 (连续响应的次数)
	FailFactor := 5.0   // 降权因子 (降权机制使用)

	// 定时的检测服务器列表中所有服务器的可用状态
	go checkServers(LB.Servers, failMax, recoverMax, FailFactor)
}

// 服务器健康检查
func checkServers(servers HttpServers, failMax int, recoverMax int, failFactor float64) {
	checker := NewHttpChecker(servers, failMax, recoverMax, failFactor)

	// 定时器，即每隔指定的时间 会向定时器的channel中发送一个时间信息
	// NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间
	ticker := time.NewTicker(3 * time.Second)
	// 服务器超时时间，即认为服务器超过该时间无响应即认为不可用
	timeout := 2 * time.Second
	for {
		select {
		case <-ticker.C:
			checker.Check(timeout)

			// 打印所有服务器的状态信息
			for _, s := range servers {
				// 计数器机制的打印信息
				fmt.Printf("server %s: %s (%d) \n", s.Host, s.Status, s.FailCount)

				// 降权机制的打印信息
				fmt.Printf("server %s: %d - %d \n", s.Host, s.Weight, s.FailWeight)
			}
			fmt.Println("-------------------------------")
		}
	}
}

// 向负载均衡的服务器列表中增添一个server
func (this *LoadBalance) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}

// 从负载均衡的服务器列表中选出一个server - 随机数算法实现负载均衡
// 随机算法存在服务器之间session同步的问题
func (this *LoadBalance) SelectServerByRand() *HttpServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.Servers))
	return this.Servers[index]
}

// 从负载均衡的服务器列表中选出一个server - 加权随机算法实现负载均衡 (增加权重)
func (this *LoadBalance) SelectServerByWeightRand() *HttpServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ServerIndexes))
	return this.Servers[ServerIndexes[index]]
}

// 从负载均衡的服务器列表中选出一个server - 加权随机算法实现负载均衡 (增加权重) ***改良版***
func (this *LoadBalance) SelectServerByWeightRand2() *HttpServer {
	sumList := make([]int, len(this.Servers))
	sum := 0
	for i := 0; i < len(this.Servers); i++ {
		sum += this.Servers[i].Weight
		sumList[i] = sum
	}
	// 比如权重是 5 1 2  => sumList[0:5, 1:6, 2:8]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(sum)
	for index, value := range sumList {
		if value > n {
			return this.Servers[index]
		}
	}
	return this.Servers[0]
}

// 从负载均衡的服务器列表中选出一个server - ip hash算法实现负载均衡
func (this *LoadBalance) SelectServerByIPHash(ip string) *HttpServer {
	// 参数ip 为客户端ip
	// 计算得到ip对应的hash值
	hashValue := crc32.ChecksumIEEE([]byte(ip))
	index := int(hashValue) % len(this.Servers)
	return this.Servers[index]
}

// 从负载均衡的服务器列表中选出一个server - 简单轮询算法实现负载均衡 + 计数器机制
func (this *LoadBalance) SelectServerByRoundRobin() *HttpServer {
	server := this.Servers[this.CurIndex]
	this.CurIndex++
	if this.CurIndex >= len(this.Servers) {
		this.CurIndex = 0
	}

	// 如果当前选到的服务器是不可用的，那么递归调用当前方法 select下一个服务器
	if server.Status == "DOWN" && !this.IsAllDown() {
		// 因为这里是递归调用，如果服务器全部不可用了，就会产生无限递归
		// 如果所有服务器全部不可用，解决办法:
		// 1. 最后一个服务器始终不设置为DOWN
		// 2. 获取server时 不再判断是否为DOWN，依然正常轮询所有服务器
		// 3. proxy直接返回一个友好的错误界面
		return this.SelectServerByRoundRobin()
	}
	return server
}

// 从负载均衡的服务器列表中选出一个server - 加权轮询算法实现负载均衡
func (this *LoadBalance) SelectServerByWeightRoundRobin() *HttpServer {
	server := this.Servers[ServerIndexes[this.CurIndex]]
	this.CurIndex++
	if this.CurIndex >= len(ServerIndexes) {
		this.CurIndex = 0
	}
	return server
}

// 从负载均衡的服务器列表中选出一个server - 加权轮询算法实现负载均衡  ***改良版 区间算法***
func (this *LoadBalance) SelectServerByWeightRoundRobin2() *HttpServer {
	server := this.Servers[0]
	sum := 0
	for i := 0; i < len(this.Servers); i++ {

		// 降权机制部分代码 start---
		realWeight := this.Servers[i].Weight - this.Servers[i].FailWeight  // 计算当前的真实权重
		if realWeight == 0 {
			// 如果真实权重为0 则跳过该服务器 (该服务器不可用)
			// 真实权重等于0 类似于 计数器机制中服务器状态字段Status等于"DOWN"
			continue
		}
		sum += realWeight
		// 降权机制部分代码 end---

		//sum += this.Servers[i].Weight
		if this.CurIndex < sum {
			server = this.Servers[i]
			if this.CurIndex == sum - 1 && i != len(this.Servers) - 1 {
				this.CurIndex++
			} else {
				this.CurIndex = (this.CurIndex + 1) % sum
			}
			break
		}
	}
	return server
}


// 从负载均衡的服务器列表中选出一个server - 平滑加权轮询算法实现负载均衡 (最重要 最常用)
func (this *LoadBalance) SelectServerBySmoothWeightRoundRobin() *HttpServer {
	// 对LoadBalance里面的servers进行排序
	sort.Sort(this.Servers)
	// 因为是降序排序，所以排序之后第一个server的当前权重是最大的
	server := this.Servers[0]

	this.Servers[0].CurWeight = this.Servers[0].CurWeight - totalWeight
	for _, s := range this.Servers {
		s.CurWeight = s.CurWeight + s.Weight
	}
	return server
}

// 从负载均衡的服务器列表中选出一个server - 平滑加权轮询算法实现负载均衡 + 降权机制
func (this *LoadBalance) SelectServerBySmoothWeightRoundRobinWithFailWeight() *HttpServer {
	// 对LoadBalance里面的servers进行排序
	sort.Sort(this.Servers)
	// 因为是降序排序，所以排序之后第一个server的当前权重是最大的
	server := this.Servers[0]

	this.Servers[0].CurWeight = this.Servers[0].CurWeight - this.GetTotalWeight()
	for _, s := range this.Servers {
		s.CurWeight = s.CurWeight + s.Weight - s.FailWeight
	}
	return server
}


// 判断服务器状态是否全部为不可用
func (this *LoadBalance) IsAllDown() bool {
	downCount := 0
	for _, server := range this.Servers {
		if server.Status == "DOWN" {
			downCount++
		}
	}
	if downCount == len(this.Servers) {
		 return true
	}
	return false
}

// 计算并获取最新的总权重 (在降权机制下 总权重是会变化的)
func (this *LoadBalance) GetTotalWeight() int {
	sum := 0
	for _, server := range this.Servers {
		// 服务器的初始权重 - 服务器当前的降权值
		realW := server.Weight - server.FailWeight
		if realW > 0 {
			sum = sum + realW
		}
	}
	return sum
}