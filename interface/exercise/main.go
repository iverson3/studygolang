package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

// 通过定义下面的三个方法 以此实现系统方法"sort.Sort(data interface)"参数所需要的接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
// Less方法决定了使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	// 先用Age字段进行排序，Age相同则用Name进行排序
	if hs[i].Age == hs[j].Age {
		return hs[i].Name < hs[j].Name
	} else {
		return hs[i].Age < hs[j].Age
	}

	// 单纯的使用Age字段进行排序
	//return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	s := []int{5, 1, 8, 4, 77, 1}
	sort.Ints(s)
	fmt.Println(s)
	
	
	var hs HeroSlice
	for i := 0; i < 6; i++ {
		hs = append(hs, Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(500),
		})
	}
	fmt.Println(hs)

	sort.Sort(hs)

	fmt.Println(hs)
}
