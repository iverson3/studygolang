package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, _ = fmt.Fprint(os.Stdout, "processing request...\n")

		select {
		case <-time.After(3 * time.Second):
			_, _ = w.Write([]byte("request processed"))
		case <-ctx.Done():
			_, _ = fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	})
	_ = http.ListenAndServe(":8888", nil)
}

func test2() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	go worker2(cancelFunc)

	breakFor := false
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("test2")
		case <-ctx.Done():
			breakFor = true
			fmt.Println("cancel from worker2 goroutine")
		}
		if breakFor {
			break
		}
	}
	fmt.Println("test2 over")
}

func worker2(cancelFunc context.CancelFunc) {
	n := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("worker2...")
		n++
		if n > 3 {
			cancelFunc()
			break
		}
	}
}

func test1()  {
	ctx := context.Background()
	typeOf := reflect.TypeOf(ctx)
	fmt.Printf("%v \n", typeOf)

	//cancelCtx, cancelFunc := context.WithCancel(ctx)
	//typeOf = reflect.TypeOf(cancelCtx)
	//fmt.Printf("%v \n", typeOf)
	//cancelFunc()

	timeoutCtx, cancelFunc2 := context.WithTimeout(ctx, 3 * time.Second)
	typeOf = reflect.TypeOf(timeoutCtx)
	//fmt.Printf("%v \n", typeOf)
	//cancelFunc2()
	go worker1(timeoutCtx)

	n := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("main")
		n++
		if n > 3 {
			break
		}
	}
	cancelFunc2()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("test")
		n++
		if n > 6 {
			break
		}
	}

	fmt.Printf("%v \n", ctx)
	fmt.Println("test1 over")
}

func worker1(ctx context.Context) {
	breakFor := false
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("work...")
		case <-ctx.Done():
			breakFor = true
			fmt.Println("cancel goroutine from main")
		}
		if breakFor {
			break
		}
	}
}
