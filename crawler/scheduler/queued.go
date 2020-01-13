package scheduler

import "studygolang/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	// workerChan 这个channel中传递的是每个准备好的worker它们自己专属的 用来收request的requestChannel  并不是worker本身
	workerChan chan chan engine.Request  // 每个worker都有自己的requestChannel (相互独立 互不干扰)
}

// 每次要创建一个新的worker时 都会调用该方法 返回一个新的requestChannel给对应的worker用来收request
func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	// 将request发给requestChannel
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// 将准备好的worker自己的requestChannel发给workerChan
	// 这个w 其实是准备好的worker传过来的它自己专属的 用来收request的requestChannel  并不是worker本身
	q.workerChan <- w
}

func (q *QueuedScheduler) Run()  {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request  // 队列里装的是 每个准备好的worker它们自己的 用来收request的requestChannel

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request  // requestChannel of active worker

			if len(requestQ) > 0 && len(workerQ) > 0 {
				// 如果request队列和worker队列都不为空 则从各自队列中取出一个
				// 在下面select中 将activeRequest发到activeWorker这个worker的requestChannel中
				// 不能在这个if里面直接发 否则可能会堵塞住
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- q.requestChan:  // 收request
				// 将Submit过来的request加入队列
				requestQ = append(requestQ, r)
			case w := <- q.workerChan:   // 收worker (worker的requestChannel)
				// 将准备好的worker放入队列
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:   // 把request发给worker (worker的requestChannel)
				// 发送成功之后则从队列中移除
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}






























