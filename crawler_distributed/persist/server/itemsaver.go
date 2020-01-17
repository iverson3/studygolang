package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/persist"
	"studygolang/crawler_distributed/rpcsupport"
)

// 自定义命令行参数
var port = flag.String("port", "", "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == "" {
		fmt.Println("must specify a port")
		return
	}

	//log.Fatal(serveRpc(elasticServerUrl, host, esIndex))
	err := serveRpc(config.ElasticServerUrl, *port, config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(elasticServerUrl, host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticServerUrl),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}


























