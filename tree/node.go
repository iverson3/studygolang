package tree

import "fmt"

// 堆 & 栈
// 局部变量分配在栈上 (函数退出 局部变量会立刻被销毁)
// 非局部变量分配在堆上 (当变量不再被使用的时候 系统垃圾回收机制会将其回收[某些语言中需要手动进行垃圾回收处理])
// go语言在这方面不用管太多，变量的定义，系统会自动判断分配到堆上还是栈上，其垃圾回收机制也会自动的进行变量回收处理

// go语言的面向对象只有封装和组合，没有继承和多态 没有构造函数
// go使用结构体 面向接口编程

// 要改变内容必须使用指针接收者
// 结构过大也考虑使用指针接收者
// 一致性： 如有指针接收者 最好都用指针接收者

// (public或private 都是相对于包package来说的)
// 变量、结构体、函数等等 首字母大写  public  (可以在其他包中进行使用)
// 变量、结构体、函数等等 首字母小写  private (在其他包中无法被使用)

// 为结构体定义的方法必须放在同一个包内，可以是不同的文件 (多个文件可以是属于同一个包)

// 定义一个结构体
type Node struct {
	Value int
	Left, Right *Node // 左右节点的指针
}

// 为结构体定义方法
func (node Node) Print()  {
	fmt.Println(node.Value)
}
// (node Node) 跟函数传参一样 都是值传递 此时在函数里面修改node的属性 是无效的，无法影响到调用该函数的对象
// 如果想要引用传递 则使用 (node *Node)
func (node *Node) SetValue(value int)  {
	// nil对象也是可以调用结构的方法的 但nil是无法访问属性的 所以下面判断如果是nil 则返回
	if node == nil {
		fmt.Println("Set value to nil node. Ignored.")
		return
	}
	node.Value = value
}

// 使用自定义的工厂函数代替构造函数 (一般不用构造函数)
func CreateNode(value int) *Node {
	// 返回了局部变量的地址，但没有任何问题 (这种做法在其他语言中可能会报错)
	return &Node{Value: value}
}

















