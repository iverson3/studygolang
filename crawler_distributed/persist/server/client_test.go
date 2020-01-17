package main

import (
	"studygolang/crawler/engine"
	"studygolang/crawler/model"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	// start up ItemSaverServer
	go serveRpc(config.ElasticServerUrl, config.RpcServeHost, config.ElasticSearchIndex)

	// sleep等待rpc服务器起来 再去连接它
	time.Sleep(time.Second)
	
	// start up ItemSaverClient
	client, err := rpcsupport.NewRpcClient(config.RpcServeHost)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Type:    "zhenai",
		Id:      "1626466343",
		Url:     "http://album.zhenai.com/u/1626466343",
		Payload: model.Profile{
			Name:       "九月",
			Gender:     "女士",
			Age:        37,
			Height:     162,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "",
			Hokou:      "",
			Xinzuo:     "",
			House:      "",
			Car:        "",
		},
	}

	// call save
	var result string
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("error: %v; result: %s", err, result)
	}
}


























