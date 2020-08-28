package main

import "fmt"

func Write(in chan int) {
	for i := 0; i < 50; i++ {
		in <- i + 1
		fmt.Println("write data to channel: ", i + 1)
	}
	close(in)
}

func Read(out chan int, flag chan bool) {
	for v := range out {
		fmt.Println("read data from channel: ", v)
	}

	//for {
	//	v, ok := <-out
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("read from channel: ", v)
	//}

	flag <- true
	close(flag)
}

func main() {
	intChan := make(chan int, 10)
	flagChan := make(chan bool)
	
	go Write(intChan)
	go Read(intChan, flagChan)

	for {
		flag := <- flagChan
		if flag {
			break
		}
	}

	for {
		_, ok := <- flagChan
		if !ok {
			break
		}
	}
}
