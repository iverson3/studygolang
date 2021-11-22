package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// select channel waitGroup once context 混合示例

// 假设有一个超长的切片，切片的元素类型是int，切片中的元素为乱序排列的。限时5s，使用多个goroutine查找切片中是否存在给定值，在找到目标值或者超时后立即结束所有goroutine的执行。
// 比如切片为：[23,32,43,76,98,54,67,32,28,61,39,.....,941,58]，查找的目标值是345，如果切片中存在目标值则输出"Found it"并且立即取消仍在执行查找任务的goroutine。
// 如果超时了则输出"Timeout，not found"，同时立即取消仍在执行查找任务的goroutine；如果查找一遍没找到目标值，则输出"Work done，but not found"并立即停止所有查找任务的goroutine。

func main() {
	maxNumber := 10000
	searchNumber := 5555
	s := generateData(maxNumber, searchNumber)
	//fmt.Printf("%v\n", s)

	search(s, searchNumber)
}

func search(s []int, targetNumber int) {
	var wg sync.WaitGroup
	var once sync.Once
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*2000)
	dataCh := make(chan int, len(s)/10)
	closeCh := make(chan struct{})
	resCh := make(chan struct{}, 1)
	workerNum := 6
	isTimeout := false

	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go worker(&wg, &once, targetNumber, dataCh, closeCh, resCh)
	}

	start := time.Now()

Exit:
	for _, v := range s {
		select {
		case <-ctx.Done():
			once.Do(func() {
				close(closeCh)
			})
			isTimeout = true
			break Exit
		case _, ok := <-closeCh:
			if !ok {
				break Exit
			}
		default:
		}
		dataCh <- v
	}
	close(dataCh)
	wg.Wait()

	duration := time.Since(start)
	fmt.Println("data len is: ", len(s))
	fmt.Printf("search time is: %v\n", duration)

	select {
	case <-resCh:
		fmt.Println("Fount it")
	default:
		if isTimeout {
			fmt.Println("Timeout, not found")
		} else {
			fmt.Println("Search work over, but not Found")
		}
	}
}

func worker(s *sync.WaitGroup, once *sync.Once, targetNumber int, dataCh chan int, closeCh chan struct{}, resCh chan struct{}) {
Exit:
	for {
		select {
		case number, ok := <-dataCh:
			if !ok {
				break Exit
			}
			if number == targetNumber {
				// 防止出现close the closed channel
				once.Do(func() {
					close(closeCh)
				})
				select {
				case resCh <- struct{}{}:
				default:
				}
				break Exit
			}
		case _, ok := <-closeCh:
			if !ok {
				break Exit
			}
		}
	}
	s.Done()
}

func generateData(maxNumber, searchNumber int) []int {
	rand.Seed(time.Now().Unix())
	lens := rand.Intn(maxNumber)
	s := make([]int, lens)

	for i := 0; i < lens; i++ {
		s[i] = rand.Intn(maxNumber)
		if s[i] == searchNumber {
			fmt.Println("search number: ", s[i])
		}
	}
	return s
}
