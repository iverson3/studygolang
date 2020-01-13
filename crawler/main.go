package main

import (
	"studygolang/crawler/engine"
	"studygolang/crawler/scheduler"
	"studygolang/crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}




























