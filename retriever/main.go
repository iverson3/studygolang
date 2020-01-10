package main

import (
	"awesomeProject1/retriever/mock"
	real2 "awesomeProject1/retriever/real"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"

// 接口的实现是隐式的, 只需要实现接口里的方法

// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口的指针
// 指针接收者实现只能以指针的方式使用, 而值接收者都可以

// 定义接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}
func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name": "stefan",
		"course": "golang",
	})
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
	//otherMethod(str string) int
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake www.baidu.com",
	})

	return s.Get(url)
	//s.otherMethod()
}


func inspect(r Retriever)  {
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

// go面向接口编程
func main() {
	//var r mock.Retriever
	//r.Contents = "somethings"

	var r Retriever

	// mock.Retriever是值接收者 所以都可以 (可取地址也可不取)
	mockRetriever1 := mock.Retriever{"this is a fake ww.baidu.com"}
	r = &mockRetriever1
	inspect(r)

	// 因为real2.Retriever是指针接收者 所以这里必须使用& 取地址
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// type assertion
	mockRetriever, ok := r.(*mock.Retriever)
	if ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	realRetriever, ok := r.(*real2.Retriever)
	if ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}

	//fmt.Println(download(r2))
	fmt.Println(session(&mockRetriever1))
}





















