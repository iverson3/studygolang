package main

import (
	"fmt"
	"sync"
	"time"
)

// go语言的传统同步机制：   WaitGroup   Mutex(互斥量)   Cond    (很少使用传统的同步机制； 都是直接使用channel进行通信)

// 传统同步机制都是通过共享内存的方式实现通信的，但因为是共享内存，在访问数据时需要对数据进行保护，否则就会出现数据访问冲突


// 自己手动实现一个原子化的int类型操作 (通过锁 解决数据访问可能出现的冲突)
// 真正使用中，最好用系统提供的原子化操作
type atomicInt struct {
	value int
	lock sync.Mutex   // 锁
}

func (a *atomicInt) increment()  {
	// 对a.value进行write操作期间 加锁，防止出现访问冲突
	a.lock.Lock()
	defer a.lock.Unlock()  // 函数执行结束 解锁
	a.value++
}

func (a *atomicInt) get() int {
	// 对a.value进行read操作期间 加锁，防止出现访问冲突
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)

	fmt.Println(a.get())
}






















