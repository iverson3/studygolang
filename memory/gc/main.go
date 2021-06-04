package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync/atomic"
	"time"
)

var (
	stop  int32
	count int64
	sum   time.Duration
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop) == 0 {
			t = time.Now()
			runtime.GC()
			sum += time.Since(t)
			count++
		}
		fmt.Printf("Gc spend avg: %v\n", time.Duration(int64(sum) / count))
	}()

	concat()
	atomic.StoreInt32(&stop, 1)
}

//func concat() {
//	limitCh := make(chan int, 800)
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < 8; i++ {
//		wg.Add(1)
//		go func() {
//			for {
//				select {
//				case _, ok := <-limitCh:
//					if !ok {
//						wg.Done()
//						return
//					}
//					s := "Go Gc"
//					s += " " + "Hello"
//					s += " " + "World"
//					_ = s
//				}
//			}
//		}()
//	}
//	for n := 0; n < 100; n++ {
//		for i := 0; i < 8; i++ {
//			limitCh<- i
//		}
//	}
//
//	close(limitCh)
//	wg.Wait()
//}

//func concat() {
//	limitCh := make(chan struct{}, 8)
//	for n := 0; n < 100; n++ {
//		for i := 0; i < 8; i++ {
//			limitCh<- struct{}{}
//			go func() {
//				s := "Go Gc"
//				s += " " + "Hello"
//				s += " " + "World"
//				_ = s
//				<-limitCh
//			}()
//		}
//	}
//}

//func concat() {
//	wg := &sync.WaitGroup{}
//	for n := 0; n < 100; n++ {
//		wg.Add(8)
//		for i := 0; i < 8; i++ {
//			go func() {
//				s := "Go Gc"
//				s += " " + "Hello"
//				s += " " + "World"
//				_ = s
//				wg.Done()
//			}()
//		}
//		wg.Wait()
//	}
//}

func concat() {
	for n := 0; n < 100; n++ {
		for i := 0; i < 8; i++ {
			go func() {
				s := "Go Gc"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
			}()
		}
	}
}