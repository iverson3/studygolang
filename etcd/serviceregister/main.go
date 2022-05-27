package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//var endpoints = []string{"1.117.77.197:20000","1.117.77.197:20002","1.117.77.197:20004"}
	var endpoints = []string{"127.0.0.1:2379"}
	registerSrv, err := NewServiceRegister(endpoints)
	if err != nil {
		log.Fatal(err)
	}

	key1 := "/web/node1"
	val1 := "localhost:8000"
	err = registerSrv.RegisterSrv(key1, val1, 5)
	if err != nil {
		log.Printf("register service to list failed, error: %v", err)
	}
	log.Println("注册服务成功， ", key1)

	time.Sleep(5 * time.Second)

	key2 := "/web/node2"
	val2 := "localhost:8001"
	err = registerSrv.RegisterSrv(key2, val2, 5)
	if err != nil {
		log.Printf("register service to list failed, error: %v", err)
	}
	log.Println("注册服务成功， ", key2)

	time.Sleep(5 * time.Second)

	err = registerSrv.RemoveSrv(key1)
	if err != nil {
		log.Printf("remove service from list failed, error: %v", err)
	}
	log.Println("移除服务成功， ", key1)

	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println("got signal: ", sig)
		done <- struct{}{}
	}()

	<-done
	// 退出之前释放资源和连接
	_ = registerSrv.Close()
	log.Println("注册服务关闭退出")
}
