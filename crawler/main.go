package main

import (
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler/persist"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
	"studygolang/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticSearchIndex)
	if err != nil {
		log.Printf("ItemSaver start up failed.")
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,      // 启动worker协程的数量
		ItemChan: itemChan,   // 这个channel负责送item数据给ItemSaver
		RequestProcessor: engine.Worker,
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




























