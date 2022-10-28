package main

import (
	"fmt"
	"sync"
)

func main() {
	var num = 5

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i+1
	}
	close(ch)

	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Println("from channel: ", v)
			}
		}()
	}

	wg.Wait()
}
