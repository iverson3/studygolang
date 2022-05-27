package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"studygolang/playground/animal"
	"sync"
	"time"
	"unsafe"
)

var (
	Token = "7f15124d79d54fb49e2ea2e632806725"
	ChannelId = "c84d31a8172324bdebaf183a5a553c3df"
)

func getSignUrl() string {
	//生成0-999之间的随机数
	timeNano := time.Now().UnixNano()
	rand.Seed(timeNano)
	nonce := fmt.Sprintf("%v", rand.Intn(1000))
	timestamp := fmt.Sprintf("%v", timeNano/1e6)

	// 排序
	params := []string{Token, timestamp, nonce}
	sort.Strings(params)

	fmt.Printf("params: %v\n", params)

	// 签名
	secret := ChannelId + Token
	signature := HMACSHA1(secret, strings.Join(params, "&"))

	fmt.Println(signature)

	v := url.Values{}
	v.Set("channel_id", ChannelId)
	v.Set("signature", signature)
	v.Set("nonce", nonce)
	v.Set("timestamp", timestamp)
	v.Set("encryption", "0")

	return v.Encode()
}

// HMACSHA1 keyStr 密钥
// value  消息内容
func HMACSHA1(key, value string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(value))
	//进行base64编码
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

type Xxx struct {
	b byte   // 1
	a int64  // 8
	e int8   // 1
	c int8   // 1
	f int32  // 4
}

func valueMutex(msg string, mutex sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(msg)
}

func doWork(ctx context.Context) {

}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	lens := len(nums)
	if lens < 3 {
		return res
	}
	sort.Ints(nums)

	var i, j, k, sum int
	for i = 0; i < lens - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}

		j, k = i + 1, lens - 1
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				for k >= 0 && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
				for j < lens && nums[j] == nums[j-1] {
					j++
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for j < lens && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k >= 0 && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}

type Test struct {
	Names []string
	n int
}

type Duck struct {
	A string   // 8+8
	B int64    // 8
	C int32    // 4 + 4 内存对齐填充4字节
}

func main() {

	dd := Duck{"duck", 3, 10}

	cPtr := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&dd)) + 16 + 8))
	//cPtr2 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&dd)) + 16 + 8))
	cPtr3 := (*[4]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&dd)) + 16 + 8))
	fillPtr := (*[4]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&dd)) + 16 + 8 + 4))

	fmt.Println(*cPtr)
	//fmt.Printf("%v %b \n", *cPtr2, *cPtr2)
	fmt.Printf("%v %b \n", *cPtr3, *cPtr3)
	fmt.Printf("%v %b \n", *fillPtr, *fillPtr)

	return
	// 访问操作结构体的私有字段
	duck := animal.NewDuck("white", 5, 10)
	//age := duck.GetAgeOfDuck()

	var offset uintptr = 16 + 8

	agePtr := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + unsafe.Sizeof(duck.Color)))
	wPtr := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset))
	//wPtr2 := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 4))

	b1 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 0))
	b2 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 1))
	b3 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 2))
	b4 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 3))
	b5 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 4))
	b6 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 5))
	b7 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 6))
	b8 := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&duck)) + offset + 7))


	fmt.Printf("%v %b \n", *b1, *b1)
	fmt.Printf("%v %b \n", *b2, *b2)
	fmt.Printf("%v %b \n", *b3, *b3)
	fmt.Printf("%v %b \n", *b4, *b4)
	fmt.Printf("%v %b \n", *b5, *b5)
	fmt.Printf("%v %b \n", *b6, *b6)
	fmt.Printf("%v %b \n", *b7, *b7)
	fmt.Printf("%v %b \n", *b8, *b8)

	// 10   1010

	// 00000000 00000000 00000000 00001010   00000000 00000000 00000000 00000000     预期存储的顺序 (大端序)
	// 00001010 00000000 00000000 00000000   00000000 00000000 00000000 00000000     预期存储的顺序 (小端序)
	// 00000000 00000000 00000000 00000000   00000000 00000000 00000000 00001010     预期输出的数据(10)对应的二进制顺序

	// 00001010 00000000 00000100 00000000   11000000 00000000 00000000 00000000

	// 00000000 00000000 00000000 11000000   00000000 00000000 00000000 00001010     实际输出的数据(824633720842)对应的二进制顺序
	// 00001010 00000000 00000000 00000000   11000000 00000000 00000000 00000000     实际存储顺序


	fmt.Println(unsafe.Sizeof(duck))    // 32
	fmt.Println(*wPtr)     // 824633720842
	//fmt.Println(*wPtr2)     //

	fmt.Println(*agePtr)
	*agePtr = 10000
	fmt.Println(*agePtr)

	//fmt.Println(age)

	return























	//
	//
	//duck := animal.NewDuck("white", 2, 10)
	//fmt.Println(duck)
	//fmt.Println(duck.Color)
	//fmt.Println(duck.Weight)
	//fmt.Println(duck.GetAgeOfDuck())
	//
	//color := *((*string)(unsafe.Pointer(uintptr(unsafe.Pointer(duck)) + unsafe.Offsetof(duck.Color))))
	//colorLen := *((*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(duck)) + unsafe.Offsetof(duck.Color) + 8)))
	//
	//age := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(duck)) + unsafe.Sizeof(duck.Color)))
	//
	//weight := *((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(duck)) + unsafe.Offsetof(duck.Weight))))
	//
	//fmt.Println(color)
	//fmt.Println(colorLen)
	//fmt.Println(*age)
	//fmt.Println(weight)
	//
	//*age = 2000
	//
	//fmt.Println(duck.GetAgeOfDuck())
	//
	//cOffset := unsafe.Offsetof(duck.Color)
	//wOffset := unsafe.Offsetof(duck.Weight)
	//fmt.Println(cOffset)
	//fmt.Println(wOffset)
	//
	//cSize := unsafe.Sizeof(duck.Color)
	//fmt.Println(cSize)
	//return
	//
	//
	//var nn int32 = 0x01020304
	//// 大端序： 00000001 00000010 00000011 00000100
	//// 小端序： 00000100 00000011 00000010 00000001      结果显示当前系统是小端序
	//
	//pointer := unsafe.Pointer(&nn)
	//ptr := uintptr(pointer)
	//
	//n1 := (*byte)(unsafe.Pointer(ptr))
	//n2 := (*byte)(unsafe.Pointer(ptr + 1))
	//n3 := (*byte)(unsafe.Pointer(ptr + 2))
	//n4 := (*byte)(unsafe.Pointer(ptr + 3))
	//
	//n5 := (*byte)(unsafe.Pointer(ptr + 4))
	//
	//fmt.Printf("%d \n", pointer)
	//fmt.Printf("0x0%x \n", *(*int32)(pointer))
	//fmt.Printf("%v %b \n", *n1, *n1)
	//fmt.Printf("%v %b \n", *n2, *n2)
	//fmt.Printf("%v %b \n", *n3, *n3)
	//fmt.Printf("%v %b \n", *n4, *n4)
	//fmt.Printf("%v \n", *n5)
	//return
	//
	//
	//
	//curBitMap := 0
	//offset := 7 - 5
	//curBitMap = curBitMap | (1 << offset)
	//
	//fmt.Println(curBitMap)
	//return



	// 0  1  2   3    4    5    6    7    8     9     10    11
	// 0  1  10  11  100  101  110  111  1000  1001  1010

	// 127   1111111
	// 119   1110111
	var days = 7
	var bit1 int64 = 119
	var bit2 = int64(math.Pow(2, float64(days))) - 1

	if bit1 & bit2 == bit2 {
		fmt.Println("bit1 全是 1")
	} else{
		fmt.Println("bit1 中含有 0")
	}

	n := bit1
	count := 0
	for n != 0 {
		count++
		n = (n - 1) & n
	}
	fmt.Printf("%d \n", count)

	fmt.Printf("%b \n", bit1)
	fmt.Printf("%b \n", bit2)
	fmt.Printf("%b \n", bit1 & bit2)

	return




	// 处理八进制表示的字符串信息
	str8 := `"\347\247\237\346\210\267\345\217\267\344\270\272\347\251\272"`
	info, _ := strconv.Unquote(str8)
	fmt.Println(info)

	return




	//threeSum([]int{0,4,0,1,0})
	//threeSum([]int{0,0,0})
	sum := threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0})
	fmt.Println(sum)
	return



	sli := []int{1, 2, 3, 4, 5}
	copy(sli[3:], sli)
	fmt.Println(sli)
	return

	str := "𠮷aߪ"

	i := len([]rune(str))
	fmt.Println(i)

	runes := []rune(str)
	fmt.Printf("%v \n", runes)
	fmt.Printf("runes[0] %s \n", string(runes[0]))
	fmt.Printf("runes[1] %s \n", string(runes[1]))


	// []byte             229        176        143          97        十进制的
	// 二进制            11100101   10110000   10001111     01100001
	// 分割标识符和有效位  1110 0101  10 110000  10 001111

	// 拼接有效的二进制位  0101 110000 001111

	// "小"对应的Unicode码  \u5c0f          十六进制的
	// 二进制            101110000001111

	bytes := []byte(str)
	fmt.Printf("%v \n", bytes)

	fmt.Printf("bytes[0]: %s \n", string(bytes[0]))
	fmt.Printf("bytes[1]: %s \n", string(bytes[1]))
	fmt.Printf("bytes[2]: %s \n", string(bytes[2]))
	fmt.Printf("bytes[3]: %s \n", string(bytes[3]))

	return


	var mm map[string]string
	mm["xx"] = "xxx"

	var sss []string
	sss[0] = "xxx"
	return



	db, err4 := sql.Open("mysql", "myscrm_rw:Yk!d3d9f3j@tcp(10.10.4.36:33306)/marketing_cloud")
	if err4 != nil {
		panic(err4)
	}

	rows, err6 := db.QueryContext(context.Background(), "select * from adv_screen")
	if err6 != nil {
		panic(err6)
	}
	_, _ = rows.Columns()
	//_ = rows.Scan()
	// defer rows.Close()

	_, err4 = os.Open("ccc")


	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	doWork(ctx)

	mu := sync.Mutex{}
	msg := "this is a message"
	valueMutex(msg, mu)

	arr := []int{1, 2, 3}
	for _, i := range arr{
		go func() {
			fmt.Println(i)
		}()
	}
	//return


	int64s := make([]int64, 0, 10)
	int64s = append(int64s, 1)
	int64s = append(int64s, 1)
	int64s = append(int64s, 1)
	int64s = append(int64s, 1)
	int64s = append(int64s, 1)

	a := "http"
	b := "www.baidu.com"
	c := a + b
	d := fmt.Sprintf("%s%s", a, b)

	slice2 := make( []string, 0)
	slice2 = append(slice2, "aaa")
	slice2 = append(slice2, "aaa")
	slice2 = append(slice2, "aaa")
	slice2 = append(slice2, "aaa")

	strings.Join(slice2, "")

	builder := strings.Builder{}
	for i := range slice2 {
		builder.Write([]byte(slice2[i]))
	}
	e := builder.String()

	fmt.Println(c,d,e)

	//struct {
	//	ptr uint64
	//	len int64
	//	cap int64
	//}

	//x := Xxx{}
	fmt.Println(unsafe.Sizeof(int64s))
	fmt.Println(&int64s[0])
	fmt.Println(&int64s[1])


	// ========================================================

	var arr3 []int
	m := sync.Mutex{}
	m.Lock()
	arr3[0] = 1
	m.Unlock()



	resp, err3 := http.Get("http://www.baidu.com")
	if err3 != nil {
		panic(err3)
	}
	//defer resp.Body.Close()
	all, err3 := ioutil.ReadAll(resp.Body)
	fmt.Println(all)



	file, err2 := os.Open("xxx")
	if err2 != nil {
		panic(err2)
	}
	readAll, err2 := ioutil.ReadAll(file)
	fmt.Println(readAll)

	_, err2 = file.Write([]byte("xxx"))
	//err2 = open.Close()

	s := "xxx"
	fmt.Printf("str: %d ", s)

	// ========================================================

	slice := []int{1, 5, 3, 8}
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= 5
	})
	fmt.Println(index)

	return
}