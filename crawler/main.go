package main

import (
	"studygolang/crawler/engine"
	"studygolang/crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}




























