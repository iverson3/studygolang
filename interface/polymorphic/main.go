package main

import "fmt"

// 多态

// 多态参数:
// 在接口的Usb接口代码案例中，方法的参数(usb Usb)既可以接收手机结构体变量，也可以接收相机结构体变量，就体现了接口的多态特性

// 多态数组：
// 通过使用接口数组，同一个数组里面可以存放不同数据类型的变量，即体现出了多态数组

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
	Price float64
}
// Phone结构体实现了Usb接口中定义的所有方法，所以称Phone结构体实现了Usb接口
func (p Phone) Start() {
	fmt.Println("phone start...")
}
func (p Phone) Stop() {
	fmt.Println("phone stop...")
}
func (p Phone) Call() {
	fmt.Println("phone call someone...")
}

type Camera struct {
	name string
}
// Camera结构体实现了Usb接口中定义的所有方法，所以称Camera结构体实现了Usb接口
func (c Camera) Start() {
	fmt.Println("camera start...")
}
func (c Camera) Stop() {
	fmt.Println("camera stop...")
}

func main() {
	// 定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	// 故体现出了多态数组的特性 (同一个数组里面可以存放不同数据类型的变量)
	var usbArr [3]Usb

	// 本来一个数组是只能存放一种数据类型的变量的，但通过接口数组，解决了这个问题；同时体现了多态数组的特性
	usbArr[0] = Phone{"小米", 1650}
	usbArr[1] = Camera{"佳能"}
	usbArr[2] = Phone{"三星", 3200}

	fmt.Println(usbArr)

	for _, v := range usbArr{
		v.Start()
		v.Stop()
		// 使用类型断言，尝试将Usb接口类型转成Phone结构体类型，返回转换结果并进行判断
		phone, ok := v.(Phone)
		if ok {
			// 转换成功则调用Phone结构体自有的方法
			phone.Call()
		}
	}
}
