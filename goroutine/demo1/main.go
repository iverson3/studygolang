package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前机器的逻辑cpu数量 (不一定是物理cpu数量)
	cpu := runtime.NumCPU()
	fmt.Println("cpu num: ", cpu)

	// 手动设置go程序最多使用多少个cpu
	//runtime.GOMAXPROCS(cpu - 1)
}
