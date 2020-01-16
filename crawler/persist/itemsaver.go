package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v6"
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler_distributed/config"
)

func ItemSaver(esIndex string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticServerUrl),
		elastic.SetSniff(false)) // 如果elasticsearch安装在docker中，必须设置sniff为 false

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			// 等待Engine那边通过返回的itemChannel发送item数据过来
			// 从itemChannel中收item 并进行存储
			item := <- out
			log.Printf("ItemSaver: got item #%d: %v", itemCount, item)

			err := Save(client, item, esIndex)
			if err != nil {
				log.Printf("ItemSaver: error saving item %v: %v", item, err)
				continue
			}
			itemCount++
		}
	}()
	return out, nil
}

// 保存数据到elasticsearch
func Save(client *elastic.Client, item engine.Item, esIndex string) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(esIndex).   // 数据库
		Type(item.Type). // 表名
		BodyJson(item)   // 数据

	if item.Id != "" {
		indexService.Id(item.Id)  // 主键Id
	}

	_, err := indexService.Do(context.Background()) // 后台运行
	if err != nil {
		return err
	}

	//fmt.Printf("%+v", resp)
	return nil
}





























