package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//test1()
	test2()
}

func test2()  {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	go func() {
		err := operation1()
		if err != nil {
			cancelFunc()
		}
	}()

	operation2(ctx)
}

func operation1() error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}
func operation2(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func test1()  {
	ctx := context.Background()
	timeoutCtx, _ := context.WithTimeout(ctx, 3 * time.Second)
	request, _ := http.NewRequest(http.MethodGet, "https://google.com", nil)
	request = request.WithContext(timeoutCtx)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("33 request failed: %v \n", err)
		return
	}

	fmt.Printf("request success! status code: %v", res.StatusCode)
}
