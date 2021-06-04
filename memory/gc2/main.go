package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// 使用 sync.Pool 复用在请求处理中需要被反复申请和释放的buf
var bufPool = sync.Pool{
	New: func() interface{} {
		//return make([]byte, 10<<20)  // 10MB
		return make([]byte, 64<<10)  // 64KB
	},
}

func main() {
	go func() {
		// 为了进行性能分析，我们还额外创建了一个监听 6060端口的 goroutine，用于使用 pprof 进行分析
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		//b := newBuf()
		b := bufPool.Get().([]byte)

		for idx := range b {
			b[idx] = 1
		}

		_, _ = fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
		bufPool.Put(b)
	})
	_ = http.ListenAndServe(":8080", nil)
}

//func newBuf() []byte {
//	return make([]byte, 10<<20)  // 10MB
//}