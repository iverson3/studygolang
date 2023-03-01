package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	bytes [][]byte
)

func main() {
	//bytes = make([]byte, 1024*1024*64 + 1024*1024 - 100)

	bytes = make([][]byte, 0, 8)

	bytes = append(bytes, make([]byte, 1024*1024*32))
	bytes = append(bytes, make([]byte, 1024*1024*1))
	bytes = append(bytes, make([]byte, 1024*123))
	bytes = append(bytes, make([]byte, 1024*901))
	bytes = append(bytes, make([]byte, 901))
	bytes = append(bytes, make([]byte, 1024*1024*3))
	bytes = append(bytes, make([]byte, 1024*1024*3))
	bytes = append(bytes, make([]byte, 1024*1024*5))

	var bm runtime.MemStats
	runtime.ReadMemStats(&bm)
	runtime.GC()
	//runtime.ReadMemStats(&em)

	// (em.Alloc - bm.Alloc)/1024

	fmt.Printf("HeapSys: %d\n", bm.HeapSys)
	fmt.Printf("Alloc: %d\n", bm.Alloc)
	fmt.Printf("HeapAlloc: %d\n", bm.HeapAlloc)
	fmt.Printf("HeapInuse: %d\n", bm.HeapInuse)
	fmt.Printf("HeapIdle: %d\n", bm.HeapIdle)
	fmt.Printf("HeapReleased: %d\n", bm.HeapReleased)

	fmt.Printf("StackSys: %d\n", bm.StackSys)
	fmt.Printf("StackInuse: %d\n", bm.StackInuse)

	fmt.Println(fmt.Sprintf("%.04f\n", float64(bm.Alloc)/1024/1024))
	//fmt.Println(fmt.Sprintf("\n%.04f\n", float64(em.Alloc)/1024/1024))

	time.Sleep(3 * time.Second)
}
