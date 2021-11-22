package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var a string

func setup(ch chan int) {
	a = "hello, world"
	ch <- 1
}

func main() {
	//os.NewFile()
	os.OpenFile("aaa.md", os.O_APPEND, 0755)
	listen, _ := net.Listen("", "")
	listen.Accept()

	ch2 := make(chan int, 1)
	go setup(ch2)

	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			ch <- true
			count++
			fmt.Printf("%d count = %v\n", j, count)
			time.Sleep(time.Millisecond)
			count--
			fmt.Printf("%d count = %v\n", j, count)
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()

	<-ch2
	fmt.Println(a)

	return

	h := NewHashMap(10)
	h.Put("aaa", 111)
	h.Put("bbb", 222)
	h.Put("aaa", 999)
	v1 := h.Get("aaa")
	v2 := h.Get("bbb")
	fmt.Println(v1)
	fmt.Println(v2)
}

func Foo(a, b int) int {
	return a + b
}

func TestQuickSort(arr []int) []int {
	quickSort(arr)
	return arr
}

func quickSort(arr []int) {
	lens := len(arr)
	if lens <= 1 {
		return
	}

	middle := arr[0]
	head, tail := 0, lens-1

	for head < tail {
		if arr[head+1] > middle { // 大于middle的元素将其交换到"最右边"
			arr[head+1], arr[tail] = arr[tail], arr[head+1]
			tail--
		} else if arr[head+1] < middle { // 小于middle的元素将其与middle进行交换
			arr[head+1], arr[head] = arr[head], arr[head+1]
			head++
		} else {
			head++
		}
	}
	// 经过上面的for循环，所有小于middle的元素都移到了左边，大于middle的元素都移到了右边

	quickSort(arr[:head])
	quickSort(arr[head+1:])
}

type HashMap struct {
	m   []*KeyPairs
	len int
	cap int
}

type KeyPairs struct {
	key   string
	value interface{}
	next  *KeyPairs
}

func NewHashMap(cap int) *HashMap {
	return &HashMap{
		m:   make([]*KeyPairs, cap, cap),
		len: 0,
		cap: cap,
	}
}

func (h *HashMap) Index(key string) int {
	return BKDRHash(key, h.cap)
}

func (h *HashMap) Put(key string, value interface{}) {
	index := h.Index(key)
	ele := h.m[index]
	if ele == nil {
		h.m[index] = &KeyPairs{
			key:   key,
			value: value,
			next:  nil,
		}
	} else {
		for {
			if ele.key == key {
				ele.value = value
				return
			}
			if ele.next == nil {
				break
			}
			ele = ele.next
		}
		ele.next = &KeyPairs{
			key:   key,
			value: value,
			next:  nil,
		}
	}

	// 扩容逻辑
}

func (h *HashMap) Get(key string) interface{} {
	index := h.Index(key)
	ele := h.m[index]
	for ele != nil {
		if ele.key == key {
			return ele.value
		}
		ele = ele.next
	}
	return nil
}

func BKDRHash(str string, cap int) int {
	seed := int(131) // 31 131 1313 13131 131313 etc..
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
}
