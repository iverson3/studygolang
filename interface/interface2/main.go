package main

import "fmt"

// 接口默认就是引用类型
// 空接口interface{} 因为没有任何方法，所以所有类型都可以算是实现了空接口，即可以把任意一个变量赋给空接口

type interfaceA interface {
	func1(int)
}

type interfaceB interface {
	func2()
}

// 同时实现上面两个接口
type Stu struct {
	name string
}
func (s Stu) func1(i int) {
	fmt.Printf("func1: %s - %d \n", s.name, i)
}
func (s Stu) func2() {
	fmt.Println("func2: ", s.name)
}


// 接口也可以继承
type interfaceX interface {
	interfaceA
	interfaceB
	func3(string) bool
}

// 此时想要实现接口interfaceX，就需要将其中继承的两个接口也一起实现
type Person struct {

}
func (p Person) func1(i int) {
	fmt.Printf("func1: %d \n", i)
}
func (p Person) func2() {
	fmt.Println("func2")
}
func (p Person) func3(s string) bool {
	fmt.Printf("func3: %s \n", s)
	return true
}


// 空接口
type T interface {

}
// 因为函数的参数类型为空接口，所以可以接受任何类型的变量作为参数
func testEmptyInterface(t T) {
	fmt.Printf("参数类型：%T； 参数值：%v \n", t, t)
}

func main() {
	s := Stu{"stefan"}
	s.func1(5)
	s.func2()

	// 一个自定义数据类型只有实现了某个接口，才能将该自定义数据类型的实例赋给接口类型变量
	var iA interfaceA = s
	var iB interfaceB = s
	iA.func1(8)
	iB.func2()


	p := Person{}
	p.func1(2)
	p.func2()
	p.func3("www")

	var iA2 interfaceA = p
	var iB2 interfaceB = p
	var iX interfaceX = p
	iA2.func1(231)
	iB2.func2()
	iX.func1(888)
	iX.func2()
	iX.func3("xxx")


	// 空接口
	var t T = p
	fmt.Println("t: ", t)

	// interface{} 即表示空接口类型
	var t2 interface{} = p
	fmt.Println("t2: ", t2)
	fmt.Printf("t2 type: %T \n", t2)

	var num float64 = 68.8
	t2 = num
	fmt.Println("t2: ", t2)
	fmt.Printf("t2 type: %T \n", t2)

	testEmptyInterface(128.55)
	testEmptyInterface(func() {})
	testEmptyInterface(p)
	var t3 interface{}
	testEmptyInterface(t3)
}
