package trap

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 随机生成指定长度的slice
func FetchData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 取slice数据中的最后一个元素
func DealData(origin []int) []int {
	start := len(origin) - 1
	tmp := origin[start:]
	return tmp
}

func DealData2(origin []int) []int {
	result := make([]int, 1)
	copy(result, origin[len(origin) - 1:])
	return result
}

// 获取并输出当前的内存使用
func PrintMem() {
	fmt.Println("====================")
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}
func GetRealSizeOf(v interface{}) (int, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}