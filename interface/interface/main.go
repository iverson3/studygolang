package main

import "fmt"

// 接口
// 任意一个变量，只要实现了某个接口中的所有方法，即表示实现了该接口
// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	
}
// Phone结构体实现了Usb接口中定义的所有方法，所以称Phone结构体实现了Usb接口
func (p *Phone) Start() {
	fmt.Println("phone start...")
}
func (p *Phone) Stop() {
	fmt.Println("phone stop...")
}

type Camera struct {
	
}
// Camera结构体实现了Usb接口中定义的所有方法，所以称Camera结构体实现了Usb接口
func (c *Camera) Start() {
	fmt.Println("camera start...")
}
func (c *Camera) Stop() {
	fmt.Println("camera stop...")
}

type Computer struct {

}
// Working()方法接收一个Usb接口类型作为参数 (接口是引用类型)
// 任何变量或结构体 只要是实现了Usb接口，就可以作为参数传入这个方法
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var comp Computer
	var p Phone
	var c Camera
	comp.Working(&p)
	comp.Working(&c)
}
