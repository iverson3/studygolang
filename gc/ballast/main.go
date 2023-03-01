package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var byteObj []byte

func main() {
	byteObj = make([]byte, 1024*1024*32)
	//runtime.KeepAlive(byteObj)

	runtime.SetFinalizer(&byteObj, func(obj *[]byte) {
		fmt.Println("obj Finalizer: ", len(*obj))
	})
	//_ = byteObj

	runtime.GC()
	fmt.Println("over")

	debug.SetGCPercent(200)   // 设置GC触发的步调比率
	debug.FreeOSMemory()

	var gcStats debug.GCStats
	debug.ReadGCStats(&gcStats)  // 读取GC的相关信息
	_ = gcStats.LastGC
	_ = gcStats.PauseTotal

	debug.PrintStack()
	debug.SetMaxStack(1024 * 1024 * 64)  // 协程栈使用的内存最多不能超过64M
	debug.SetMaxThreads(2)

	debug.SetTraceback("1")

	f, _ := os.Open("xxx.log")
	debug.WriteHeapDump(f.Fd())
	stat, _ := f.Stat()

	r, w, _ := os.Pipe()
	w.Write([]byte("xxx"))
	go func() {
		r.Read([]byte(""))
	}()

	f2, _ := os.Open("xxx.log")
	stat2, _ := f2.Stat()
	isSame := os.SameFile(stat2, stat)
	if isSame {
		log.SetOutput(os.Stdout)
		log.Print("f2 is same with f")
	}

	dirList, _ := ioutil.ReadDir("dirname")
	for _, dir := range dirList {
		dir.IsDir()
		dir.Name()
	}

	nopCloser := ioutil.NopCloser(f)
	nopCloser.Read([]byte("xxx"))

	runtime.LockOSThread()
	runtime.UnlockOSThread()

	debug.ReadBuildInfo()

	time.Sleep(2 * time.Second)
	return

	var upCounter int
	// 输入
	var summonRate float64 = 1

	fmt.Println("Official Rate (%): ")
	fmt.Print("Summon times        :  ")
	start := time.Now()

	for i := 50; i < 600; i += 50 {
		fmt.Printf("  %d   |", i)
	}
	fmt.Println()

	for k := 0; k < 7; k++ {
		fmt.Printf("more than %d star(s) : ", k)
		for j := 50; j < 600; j += 50 {
			upCounter = 0
			for i :=0; i < 100000; i++ {
				if upSummon(summonRate, j) > k {
					upCounter++
				}
			}
			fmt.Printf(" %.02f%s |", float64(upCounter)/1000, "%")
		}
		fmt.Println()
	}
	fmt.Printf("took: %.02f s\n", time.Since(start).Seconds())
	time.Sleep(3 * time.Second)
}

func upSummon(summonRate float64, summonTimes int) int {
	var upCounter int
	var lastSummon int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < summonTimes; i++ {
		lastSummon++
		if lastSummon == 100 && rand.Float64() * 15 < summonRate {
			upCounter++
			lastSummon = 0
		} else if rand.Float64() * 1000 < summonRate * 10 {
			upCounter++
		}
	}
	return upCounter
}

func test() {
	counter := make([]int, 5)

	for i := 0; i < 5; i++ {
		summon := upSummon(1, 1)

		for j := 0; j < 7; j++ {
			if summon <= j {
				break
			}
			counter[j]++
		}
	}
}
