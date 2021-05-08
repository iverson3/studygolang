package main

import "fmt"

// Go 程序会在 2 个地方为变量分配内存，一个是全局的堆(heap)空间用来动态分配内存，另一个是每个 goroutine 的栈(stack)空间
// Go 语言实现垃圾回收(Garbage Collector)机制，因此呢，Go 语言的内存管理是自动的，通常开发者并不需要关心内存分配在栈上，还是堆上。
// 但是从性能的角度出发，在栈上分配内存和在堆上分配内存，性能差异是非常大的。

//在函数中申请一个对象，如果分配在栈中，函数执行结束时自动回收，如果分配在堆中，则在函数结束后某个时间点进行垃圾回收。
//在栈上分配和回收内存的开销很低，只需要 2 个 CPU 指令：PUSH 和 POP，一个是将数据 push 到栈空间以完成分配，pop 则是释放空间，也就是说在栈上分配内存，消耗的仅是将数据拷贝到内存的时间
//在堆上分配内存，一个很大的额外开销则是垃圾回收。Go 语言使用的是标记清除算法，并且在此基础上使用了三色标记法和写屏障技术，提高了效率;

//标记清除收集器是跟踪式垃圾收集器，其执行过程可以分成标记（Mark）和清除（Sweep）两个阶段：
// 标记阶段 — 从根对象出发查找并标记堆中所有存活的对象；
// 清除阶段 — 遍历堆中的全部对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表。
// 标记清除算法的一个典型耗时是在标记期间，需要暂停程序（Stop the world，STW），标记结束之后，用户程序才可以继续执行。



// 逃逸分析 (函数内部的变量原本是要分配在stack上的，如果因为某些原因最终分配到了heap上，这个现象就被称之为逃逸)

// 在 C 语言中，可以使用 malloc 和 free 手动在堆上分配和回收内存。Go 语言中，堆内存是通过垃圾回收机制自动管理的，无需开发者指定。
//那么，Go 编译器怎么知道某个变量需要分配在栈上，还是堆上呢？编译器决定内存分配位置的方式，就称之为逃逸分析(escape analysis)。逃逸分析由编译器完成，作用于编译阶段。

// 指针逃逸
// interface{}动态类型逃逸
// 栈空间不足导致的逃逸
// 因闭包而产生的逃逸

type Demo struct {
	name string
	age int
}

func main() {
	//t1()
	//t2()
	//t3()
	//t4()
	t5()
}

// 指针逃逸
//func t1() {
//	demo := createDemo("xxx", 10)
//	fmt.Println(demo)
//}
//func createDemo(name string, age int) *Demo {
//	dd := &Demo{name: name, age: age}
//	return dd
//}


// interface{}动态类型逃逸
//func t2() {
//	val := "ccc"
//	printInterfaceValue(val)
//}
//// 如果将interface{}动态类型换成某个确定的类型比如string，val变量就不会发生逃逸了 (不过value变量还是会发生逃逸，因为Println函数参数就是interface{}类型)
//func printInterfaceValue(value interface{}) {
//	fmt.Println(value)
//}


// 栈空间不足导致的逃逸
// 结论：当切片占用内存超过一定大小，或无法确定当前切片长度时，对象占用内存将在堆上分配
func t3() {
	// 操作系统对内核线程使用的栈空间是有大小限制的，64 位系统上通常是 8 MB。可以使用 ulimit -a 命令查看机器上栈允许占用的内存的大小
	// 对于 Go语言来说，运行时(runtime) 尝试在 goroutine需要的时候动态地分配栈空间，goroutine的初始栈大小为 2KB。
	// 当 goroutine被调度时，会绑定内核线程执行，栈空间大小最终也不会超过操作系统的限制。

	// 对 Go编译器而言，超过一定大小的局部变量将逃逸到堆上，不同的Go版本的大小限制可能不一样
	//generate8191()
	//generate8192()
	//generate(1)
}
//func generate8191() {
//	nums := make([]int, 8191) // < 64KB
//	for i := 0; i < 8191; i++ {
//		nums[i] = rand.Int()
//	}
//}
//func generate8192() {
//	nums := make([]int, 8192) // = 8192 * 8 = 64KB
//	for i := 0; i < 8192; i++ {
//		nums[i] = rand.Int()
//	}
//}
//func generate(n int) {
//	nums := make([]int, n) // 不确定大小
//	for i := 0; i < n; i++ {
//		nums[i] = rand.Int()
//	}
//}


// 因闭包而产生的逃逸
//func t4() {
//	incr := createFunc()
//	fmt.Println(incr())
//	fmt.Println(incr())
//}
//// createFunc()返回值是一个闭包函数，该闭包函数访问了外部变量 n，那变量 n 将会一直存在，直到 incr 被销毁。
//// 很显然，变量 n 占用的内存不能随着函数 createFunc()的退出而回收，因此将会逃逸到堆上
//func createFunc() func() int {
//	n := 0
//	return func() int {
//		n++
//		return n
//	}
//}


// slice在append时元素个数超过其cap，重新分配内存 存放新的底层数组
// 结论：好像并不会发生逃逸
func t5()  {
	ss := make([]int, 2, 2)
	ss[0] = 1
	ss[1] = 2
	//fmt.Println(cap(ss))   // cap is 2
	ss = appendEleToSlice(ss)
	fmt.Println(ss[3])

	//changeSliceValue(ss)
	//fmt.Println(ss[0])
}
func appendEleToSlice(s []int) []int {
	//fmt.Println(cap(s))  // cap is 2
	s = append(s, 3)
	s = append(s, 4)
	//fmt.Println(cap(s))  // cap is 4
	return s
}
//func changeSliceValue(s []int) {
//	// 有了这行代码，之后对slice的修改就无法影响到函数外的原slice变量了
//	// 因为append使slice的元素个数超过了当前的cap，所以会重新分配内存 用来存放新的更大的底层数组，并且s变量会指向新的底层数组，而函数外的ss还是指向原来的底层数组
//	s = append(s, 3)
//	s[0] = 100
//}