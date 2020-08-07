package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan int)

	for j := 0; j < 3; j++ {
		go newClient(out, j)
		time.Sleep(time.Millisecond * 200)
	}
	for j := 0; j < 3; j++ {
		out <- j
	}

	time.Sleep(time.Millisecond * 2000)

	i := 3
	defer func() {
		fmt.Printf("this is from defer:%d\n", i)
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Recover Error: %v", err)
		} else {
			panic(r)
		}
	}()

	i = i * 2 - 6
	res := 1 / i
	fmt.Println(res)
}

func newClient(in chan int, no int)  {
	for {
		fmt.Printf("number: %v, time: %v\n", no, time.Now().UnixNano() / 1e6)
		time.Sleep(time.Millisecond * 1000)
		num := <-in
		fmt.Println(no, num)
	}
}
