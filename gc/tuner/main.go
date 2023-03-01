package main

import (
	"fmt"
	"github.com/cch123/gogctuner/gotuner"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"

	"runtime"
	"time"
)

//var byteObj []byte

var glob [][]byte

func main() {
	//byteObj = make([]byte, 1024*1024*100)

	glob = make([][]byte, 3)

	log.SetOutput(os.Stdout)
	go gotuner.NewTuner2(false, 0.8, gotuner.SetLogger(&gotuner.StdLoggerAdapter{DebugEnabled: true}))

	runtime.GC()

	obj1 := make([]byte, 1024*1024*512)
	glob[0] = obj1
	runtime.GC()
	obj2 := make([]byte, 1024*1024*1024)
	glob[1] = obj2
	runtime.GC()
	obj3 := make([]byte, 1024*1024*1024*2)
	glob[2] = obj3
	runtime.GC()

	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		panic(err)
	}
	info, err := p.MemoryInfo()
	if err != nil {
		panic(err)
	}

	percent, err := p.MemoryPercent()
	if err != nil {
		panic(err)
	}
	fmt.Println("percent: ", percent)

	fmt.Println("VMS: ", info.VMS)
	fmt.Println("RSS: ", info.RSS)
	//fmt.Println("HWM: ", info.HWM)
	//fmt.Println("Swap: ", info.Swap)

	time.Sleep(5 * time.Second)
}

