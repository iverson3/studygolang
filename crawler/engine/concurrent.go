package engine

import (
	"github.com/go-redis/redis"
	"log"
	"studygolang/crawler_distributed/duplicate"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
	RedisClientChan chan *redis.Client
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)  // 通过channel并发的分发request (给worker) [注意： 一定要并发的分发(使用goroutine) 如果用同步的方式分发request  可能会导致循环等待]
	WorkerChan() chan Request
	Run()
}

// 单独定义成一个interface  然后通过组合的方式放入Scheduler
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 循环创建多个worker  每个worker都是一个goroutine
	for i := 0; i < e.WorkerCount; i++ {
		// 从scheduler获取当前马上要创建的worker自己专属的requestChannel
		in := e.Scheduler.WorkerChan()
		// 每个worker有自己的requestChannel  用来收request
		e.createWorker(in, out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		// 从outChannel中收"请求的结果"
		result := <- out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) == false {
				e.Scheduler.Submit(request)
			}
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	go func() {
		for {
			// 代码解释： 这里当前这个worker准备好接受request之后 就调用WorkerReady()告诉scheduler自己已准备好 并将自己的requestChannel给它
			// 这样 scheduler在有了request之后 就会把request发到这个worker自己的requestChannel中 下面代码就是从自己的channel中收request
			ready.WorkerReady(in)

			// 从自己的requestChannel中收request
			request := <- in

			// redis去重
			exist := duplicate.IsDuplicate(e.RedisClientChan, request.Url)
			if exist {
				continue
			}

			result, err := e.RequestProcessor(request)
			if err != nil {
				log.Printf("worker(Processor) error: %v", err)
				continue
			}
			// 向outChannel中发"请求的结果"
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

// url去重
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}






























