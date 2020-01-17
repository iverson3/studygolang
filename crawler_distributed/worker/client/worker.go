package client

import (
	"net/rpc"
	"studygolang/crawler/engine"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		var result worker.ParseResult
		client := <- clientChan
		err := client.Call(config.CrawlServiceRpc, worker.SerializeRequest(req), &result)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(result), nil
	}
}






































