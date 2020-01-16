package main

import (
	"gopkg.in/olivere/elastic.v6"
	"studygolang/crawler_distributed/config"
	"studygolang/crawler_distributed/persist"
	"studygolang/crawler_distributed/rpcsupport"
)

func main() {
	//log.Fatal(serveRpc(elasticServerUrl, host, esIndex))
	err := serveRpc(config.ElasticServerUrl, config.RpcServeHost, config.ElasticSearchIndex)
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


























