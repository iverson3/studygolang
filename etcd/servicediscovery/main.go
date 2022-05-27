package main

import (
	"log"
	"time"
)

func main() {
	//var endpoints = []string{"1.117.77.197:20000","1.117.77.197:20002","1.117.77.197:20004"}
	var endpoints = []string{"127.0.0.1:2379"}
	srv := NewServiceDiscovery(endpoints)
	defer func() {
		srv.Close()
		log.Println("发现服务关闭退出")
	}()

	prefix := "/web/"
	err := srv.WatchService(prefix)
	if err != nil {
		log.Printf("watch prefix: %s failed, error: %v", prefix, err)
	}
	//prefix2 := "/gRPC/"
	//err = srv.WatchService(prefix2)
	//if err != nil {
	//	log.Printf("watch prefix: %s failed, error: %v", prefix2, err)
	//}

	for {
		select {
		case <-time.Tick(4 * time.Second):
			log.Println(srv.GetServices())
		}
	}
}
