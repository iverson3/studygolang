package client

import (
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			// 等待Engine那边通过返回的itemChannel发送item数据过来
			// 从itemChannel中收item 然后调用RpcServe进行数据存储 (ItemSaverService.Save)
			item := <- out
			log.Printf("ItemSaver: got item #%d: %v", itemCount, item)

			var result string
			// call rpcServe to save item
			err = client.Call(config.ItemSaverRpc, item, &result)

			if err != nil || result != "ok" {
				log.Printf("ItemSaver: error saving item %v: %v", item, err)
				continue
			}
			itemCount++
		}
	}()
	return out, nil
}