package main

import (
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(config.RpcServeHost)
	if err != nil {
		log.Printf("ItemSaver start up failed.")
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,      // 启动worker协程的数量
		ItemChan: itemChan,   // 这个channel负责送item数据给ItemSaver
	}
	//e.Run(engine.Request{
	//	Url:        config.SeedUrl,
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}




























