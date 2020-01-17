package main

import (
	"flag"
	"fmt"
	"studygolang/crawler_distributed/rpcsupport"
	"studygolang/crawler_distributed/worker"
)

// 自定义命令行参数
var port = flag.String("port", "", "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == "" {
		fmt.Println("must specify a port")
		return
	}

	err := rpcsupport.ServeRpc(*port, worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}

























