package main

import (
	"fmt"
)

// 哪些数据类型可以作为map的key
// map使用哈希表 必须可以比较相等
// 除了slice map function的内建类型都可以作为key
// Struct类型不包含上述字段，也可以作为key

// map 数据类型
func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)  // m2 == empty map
	var m3 map[string]int       // m3 == nil

	fmt.Println(m, m2, m3)


	// 遍历map (遍历的顺序是不确定的 因为map本身就是无序的)
	// 如需排序，则需手动对key进行排序  将key放入slice然后排序(slice是支持排序的)
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 获取map长度
	mLen := len(m)
	fmt.Println(mLen)

	// 获取value
	courseName, err := m["course"]
	fmt.Println(courseName, err)  // err == true

	courseName111, err := m["course111"]  // key不存在
	if err {
		fmt.Println(courseName111)
	} else {
		fmt.Println("key does not exist")
	}

	// delete elements from map
	name, err := m["name"]
	fmt.Println(name, err)

	delete(m, "name")

	name, err = m["name"]
	fmt.Println(name, err)
}
