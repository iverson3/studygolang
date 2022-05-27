package main

import (
	"context"
	"errors"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

// ServiceRegister 服务注册及健康检查
type ServiceRegister struct {
	cli *clientv3.Client
	mu sync.Mutex
	// 所有注册的服务列表
	srvList map[string]*Service
}

type Service struct {
	leaseID clientv3.LeaseID  // 租约ID
	key string
	val string
	// 租约keepalive对应的chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

func NewServiceRegister(endpoints []string) (*ServiceRegister, error)  {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	service := &ServiceRegister{
		cli: client,
		srvList: make(map[string]*Service),
	}

	return service, nil
}

// RegisterSrv 注册服务
func (s *ServiceRegister) RegisterSrv(key, val string, lease int64) error {
	srv := &Service{
		key: key,
		val: val,
	}
	s.mu.Lock()
	s.srvList[key] = srv
	s.mu.Unlock()

	err := s.putKeyWithLease(key, lease)
	if err != nil {
		return err
	}
	// 开始监听续约情况
	go s.listenLeaseRespChan(key)
	return nil
}

// 设置租约
func (s *ServiceRegister) putKeyWithLease(key string, lease int64) error {
	// 设置租约时间
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	resp, err := s.cli.Grant(ctx, lease)
	cancelFunc()
	if err != nil {
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(context.Background(), key, s.srvList[key].val, clientv3.WithLease(resp.ID))
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
		return err
	}
	// 设置租约，定期发送请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return err
	}
	s.srvList[key].leaseID = resp.ID
	s.srvList[key].keepAliveChan = leaseRespChan
	return nil
}

// ListenLeaseRespChan 监听续约情况
func (s *ServiceRegister) listenLeaseRespChan(key string) {
	for leaseKeepResp := range s.srvList[key].keepAliveChan {
		log.Printf("续约成功, key: %s, leaseKeepResp: %v \n", key, leaseKeepResp)
	}
	log.Println("关闭续约, key: ", key)
}

// RemoveSrv 从服务列表中移除指定的服务
func (s *ServiceRegister) RemoveSrv(key string) error {
	srv, ok := s.srvList[key]
	if !ok {
		return errors.New("服务不存在")
	}
	// 撤销对应服务的租约
	_, err := s.cli.Revoke(context.Background(), srv.leaseID)
	if err != nil {
		return err
	}

	s.mu.Lock()
	delete(s.srvList, key)
	s.mu.Unlock()
	return nil
}

// Close 关闭注册服务
func (s *ServiceRegister) Close() error {
	// 遍历撤销所有服务的租约
	for _, srv := range s.srvList {
		_, err := s.cli.Revoke(context.Background(), srv.leaseID)
		if err != nil {
			return err
		}
	}

	log.Println("撤销所有服务的租约")
	return s.cli.Close()
}




















