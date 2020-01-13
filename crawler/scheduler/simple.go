package scheduler

import "studygolang/crawler/engine"

type SimpleScheduler struct {
	// 所有的worker都从这个channel中收request
	workerChan chan engine.Request   // 所有的worker共用一个requestChannel
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	// 给所有的worker返回的都是同一个requestChannel
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	// 创建一个requestChannel给所有worker共用
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 使用goroutine并发的分发request
	// 这里必须使用goroutine向workerchannel中发request 否则会导致循环等待
	// 循环等待： 向channel中发数据 和 从channel中收数据 彼此等待阻塞，导致卡死
	go func() { s.workerChan <- r }()
}


