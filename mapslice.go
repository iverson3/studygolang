package main

import (
	"fmt"
	"sort"
)

type student struct {
	name string
	age int
	score float64
	pass bool
}

func main() {
	// map切片

	heros := make([]map[string]string, 2)

	if heros[0] == nil {
		heros[0] = make(map[string]string)
	}
	heros[0]["name"] = "aaa"
	heros[0]["age"] = "111"

	if heros[1] == nil {
		heros[1] = make(map[string]string)
	}
	heros[1]["name"] = "bbb"
	heros[1]["age"] = "222"

	// 超出切片长度之后，添加新元素需要使用append()函数进行添加
	newHero := make(map[string]string)
	newHero["name"] = "ccc"
	newHero["age"] = "333"
	heros = append(heros, newHero)

	fmt.Printf("heros = %v", heros)
	fmt.Println()



	// 对map进行排序
	// map本身是无序的，想要对它进行排序，则需要先对map的key进行排序，再根据排序的key去取对应的map值
	scores := make(map[string]int, 2)
	scores["james"] = 99
	scores["kobe"] = 98
	scores["jordan"] = 100

	var keys []string
	for k, _ := range scores{
		keys = append(keys, k)
	}

	fmt.Println("keys排序前: ", keys)
	sort.Strings(keys)
	fmt.Println("keys排序后: ", keys)

	for _, v := range keys{
		fmt.Println(scores[v])
	}



	// 最常充当map的value的数据类型是struct结构体
	students := make(map[int]student)

	students[10001] = student{
		name:  "stefan",
		age:   26,
		score: 98,
		pass:  true,
	}
	students[10002] = student{
		name:  "jake",
		age:   25,
		score: 88,
		pass:  false,
	}

	fmt.Println("studens: ", students)

}
