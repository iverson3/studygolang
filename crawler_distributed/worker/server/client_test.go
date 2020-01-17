package main

import (
	"fmt"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/rpcsupport"
	"studygolang/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	//req := worker.Request{
	//	Url:    "https://album.zhenai.com/u/1626466343",
	//	Parser: worker.SerializedParser{
	//		Name: config.ParseProfile,
	//		Args: parser.ProfileParser{UserName: "九月", Sex: "女士"},
	//	},
	//}
	req := worker.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: worker.SerializedParser{
			Name: config.ParseCity,
			Args: nil,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}






















