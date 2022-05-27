package main

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli *clientv3.Client
	// 服务列表
	serverList map[string]string
	lock sync.Mutex
}

func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceDiscovery{
		cli:        client,
		serverList: make(map[string]string),
	}
}

func (s *ServiceDiscovery) WatchService(prefix string) error {
	// 根据前缀获取现有的key
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 3*time.Second)
	resp, err := s.cli.Get(ctx, prefix, clientv3.WithPrefix())
	cancelFunc()
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}

	// 监视前缀，修改变更的server
	go s.watcher(prefix)
	return nil
}

func (s *ServiceDiscovery) watcher(prefix string) {
	watchChan := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix: %s .....", prefix)
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: // 新增或修改
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: // 删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = val
	log.Printf("put key: %s, val: %s", key, val)
}

func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Printf("delete key: %s", key)
}

func (s *ServiceDiscovery) GetServices() []string {
	addrs := make([]string, 0)
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, addr := range s.serverList {
		addrs = append(addrs, addr)
	}
	return addrs
}

func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}
