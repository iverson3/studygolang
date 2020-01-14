package main

import (
	"studygolang/crawler/engine"
	"studygolang/crawler/persist"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"
const cityUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
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




























