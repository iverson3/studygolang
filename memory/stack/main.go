package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

func main() {
	//s1 := "xxx"
	//f1(&s1)

	sli := FetchData(128 * 1024)
	s2 := strings.Join(sli, "")

	//runtime.GC()

	PrintMem()
	f2(&s2)
}

//func f1(s *string) {
//	fmt.Println(s)
//}

func f2(s *string) {
	PrintMem()
	fmt.Println((*s)[:2])
	fmt.Println(len(*s))

	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(*s))

	size, _ := GetRealSizeOf(*s)
	fmt.Println(size)
}


func FetchData(n int) []string {
	rand.Seed(time.Now().UnixNano())
	nums := make([]string, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, string(rand.Int()))
	}
	return nums
}
// 获取并输出当前的内存使用
func PrintMem() {
	fmt.Println("====================")
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%.2f MB\n", float64(rtm.Alloc)/1024./1024.)
}
func GetRealSizeOf(v interface{}) (int, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}