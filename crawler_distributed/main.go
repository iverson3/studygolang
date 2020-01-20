package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
	"studygolang/crawler/engine"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/duplicate"
	itemSaverClient "studygolang/crawler_distributed/persist/client"
	"studygolang/crawler_distributed/rpcsupport"
	workerClient "studygolang/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
	workerCount = flag.Int("worker_count", 0, "count of workers")
)

func main() {
	flag.Parse()
	if *itemSaverHost == "" || *workerHosts == "" || *workerCount == 0 {
		fmt.Println("missing parameters; use --help to see parameters")
		return
	}

	itemChan, err := itemSaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		log.Printf("ItemSaver start up failed.")
		panic(err)
	}

	redisClientChan, count := duplicate.CreateRedisClientPool(config.RedisServerUrl, 5)
	if count == 0 {
		panic("Can not connected to redis server.")
	}

	hostSlice := strings.Split(*workerHosts, ",")
	pool := createClientPool(hostSlice)
	processor := workerClient.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: *workerCount,    // 启动worker协程的数量
		ItemChan: itemChan,           // 这个channel负责送item数据给ItemSaver
		RequestProcessor: processor,
		RedisClientChan: redisClientChan,
	}

	//e.Run(engine.Request{
	//	Url:    config.SeedUrl,
	//	Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	//})
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewRpcClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s.", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}



























