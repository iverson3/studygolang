package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func computeCapacity(param int) (size int) {
	if param <= 16 {
		return 16
	}
	n := param - 1
	fmt.Println(n)
	n |= n >> 1
	fmt.Println(n)
	n |= n >> 2
	fmt.Println(n)
	n |= n >> 4
	fmt.Println(n)
	n |= n >> 8
	fmt.Println(n)
	n |= n >> 16
	fmt.Println(n)
	if n < 0 {
		return math.MaxInt32
	}
	return n + 1
}

func main() {
	url := "https://www.baidu.com/aaa/bbb/index.jpeg"

	s := replaceUrl(url, "www.google.com")
	fmt.Println(s)


	return
	// 1433          pass
	// 661543647     pass
	// 6615437765    pass
	res := nextMaxNum2(1234)
	fmt.Println(res)
	return

	//b := test(1024 * 1024 * 100)
	//c := b
	//fmt.Println(c[0])
	//b = nil
	//c = nil

	//socketFd, _ := syscall.Socket(syscall.BASE_PROTOCOL, syscall.AF_INET6, syscall.TCP_NODELAY)
	//syscall.Bind(socketFd, syscall.Sockaddr())


	fd, _ := syscall.Open("./a.txt", syscall.O_SYNC|syscall.O_APPEND, 1)
	newSeek, _ := syscall.Seek(fd, 10, 0)
	fmt.Println(newSeek)

	buf := make([]byte, 10)
	n, err := syscall.Read(fd, buf)
	fmt.Println(n, err)

	fmt.Println(string(buf))

	//return

	writeBuf := []byte("12345678")
	n, err = syscall.Write(fd, writeBuf)
	fmt.Println(n, err)
	err = syscall.Fsync(fd)
	fmt.Println(err)

	PrintMemUsage()

	return

	gogc := os.Getenv("GOGC")
	path := os.Getenv("GOPATH")
	pageSize := os.Getpagesize()
	fmt.Println(gogc)
	fmt.Println(path)
	fmt.Println(pageSize)

	//debug.SetGCPercent(200)
	//debug.FreeOSMemory()

	runtime.GC()

	stat := &debug.GCStats{}
	debug.ReadGCStats(stat)

	fmt.Println(stat)

	//debug.SetMaxStack(1024 * 1024 * 32)
	//debug.SetMaxThreads(10)
	return

	var shardCount int = 27

	capacity := computeCapacity(shardCount)

	fmt.Println(shardCount, capacity)

	return
}

func test(size int) []byte {
	a := make([]byte, size)
	for i := 0; i < size; i++ {
		a[i] = 'x'
	}
	return a
}


// 以下是打印内存监控数据的工具函数，与业务逻辑无关
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v", m.NumGC)
	fmt.Printf("\tAllocObjCnt = %v", m.Mallocs)
	fmt.Printf("\tSTW = %v\n", m.PauseTotalNs)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func nextMaxNum(num int64) int64 {
	numStr := strconv.FormatInt(num, 10)
	max := 0
	for i := 0; i < len(numStr) - 1; i++ {
		if numStr[i+1] > numStr[i] {
			max = i + 1
		}
	}

	if max == 0 {
		return num
	}
	newNumStr := numStr[:max-1] + string(numStr[max]) + string(numStr[max-1]) + numStr[max + 1:]

	res, _ := strconv.ParseInt(newNumStr, 10, 64)

	return res
}

// 6615437765   len=9
// 0123456789

func nextMaxNum2(num int64) int64 {
	str := strconv.FormatInt(num, 10)
	pos := 0
	var res int64
	passNums := make([]int, 0, len(str))

	for i := len(str) - 1; i > 0; i-- {
		if str[i] > str[i-1] {
			pos = i
			break
		} else {
			int1, _ := strconv.ParseInt(string(str[i]), 10, 32)
			passNums = append(passNums, int(int1))
		}
	}

	if len(passNums) > 1 {
		sort.Ints(passNums)
	}

	if len(passNums) == 0 {
		newNumStr := str[:pos-1] + string(str[pos]) + string(str[pos-1]) + str[pos + 1:]
		res, _ = strconv.ParseInt(newNumStr, 10, 64)
	} else {
		min := passNums[0]
		newNumStr := str[:pos-1] + strconv.Itoa(min)

		// 把 pos 和 pos-1 两位数字 放入切片
		int2, _ := strconv.ParseInt(string(str[pos-1]), 10, 32)
		int3, _ := strconv.ParseInt(string(str[pos]), 10, 32)
		passNums[0] = int(int2)
		passNums = append(passNums, int(int3))
		sort.Ints(passNums)

		for i := 0; i < len(passNums); i++ {
			newNumStr = newNumStr + strconv.Itoa(passNums[i])
		}
		res, _ = strconv.ParseInt(newNumStr, 10, 64)
	}
	return res
}

func replaceUrl(url string, target string) string {
	pattern := `^(http|https)://([^/]+)/.*\.(jpeg|jpg|png)$`

	var reg = regexp.MustCompile(pattern)
	submatch := reg.FindStringSubmatch(url)
	fmt.Println(submatch)
	replace := strings.Replace(url, submatch[2], target, 1)
	return replace
}