package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

const (
	elasticServerUrl = "http://47.107.149.234:9200"
	eIndex = "dating_profile"
	eType = "zhenai")


func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			// 等待Engine那边通过返回的itemChannel发送item数据过来
			// 从itemChannel中收item 并进行存储
			item := <- out
			log.Printf("ItemSaver: got item #%d: %v", itemCount, item)

			_, err := save(item)
			if err != nil {
				log.Printf("ItemSaver: error saving item %v: %v", item, err)
				continue
			}
			itemCount++
		}
	}()
	return out
}

// 保存数据到elasticsearch
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticServerUrl),
		elastic.SetSniff(false)) // 如果elasticsearch安装在docker中，必须设置sniff为 false

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index(eIndex). // 数据库
		Type(eType). // 表名
		BodyJson(item). // 数据 (不设置id  让系统自动生成)
		Do(context.Background()) // 后台运行

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v", resp)
	return resp.Id, nil
}





























