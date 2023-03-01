package main

import (
	"fmt"
	"runtime"
	"time"
)

// GC优化技术之 ballast

//var (
//	ballast = flag.Bool("ballast", false, "run program with Ballast")
//)

func main() {
	ballastObj := make([]byte, 2*1024*1024*1024)
	runtime.KeepAlive(ballastObj)

	//debug.SetGCPercent(100)

	//flag.Parse()
	//withBallast := *ballast
	//
	//fmt.Println(withBallast)
	//if withBallast {
	//	//WithBallast()
	//}

	start := time.Now()

	for i := 0; i < 10; i++ {
		trees := allocMem()

		b := (*trees)[0].B
		_ = b
	}

	fmt.Printf("\ntook %.02f s\n", float64(time.Since(start).Milliseconds())/1000)
}

type Tree struct {
	B []byte
}

func allocMem() *[]*Tree {
	trees := make([]*Tree, 50)
	for i := 0; i < 50; i++ {
		trees[i] = &Tree{B: make([]byte, 1024*1024*10)}
		//trees[i].B[10*1024*1024 - 1] = 'a'
	}
	return &trees
}

//func WithBallast() {
//	// 声明一个10G的字节切片(但不用它)，占着堆内存使用空间（不占用实际内存使用），拉高GC触发的阈值
//	//ballast := make([]byte, 10*1024*1024*1024)
//	ballast := make([]byte, 1024*1024*1024)
//
//	// 确保ballast不会被GC给回收掉
//	runtime.KeepAlive(ballast)
//}
