package main

import (
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler/persist"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"
const cityUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
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
	//	Url:        url,
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        cityUrl,
		ParserFunc: parser.ParseCity,
	})
}




























