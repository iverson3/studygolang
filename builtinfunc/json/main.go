package main

import (
	"encoding/json"
	"fmt"
)

// json的序列化和反序列化

// 在golang中，字符串 数字 数组 slice map 结构体等，都可以进行json序列化

// json序列化：将具有key-value结构的数据类型(比如struct map slice) 序列化成json字符串
// json反序列化：将json字符串还原成其序列化之前的数据类型

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Height float64 `json:"height"`
	Skill *[]string `json:"skill"`
	FamilyShip *map[string]interface{} `json:"family_ship"`
}

func testStructJson()  {
	skills := []string{"跳舞", "钢琴"}
	familyShips := make(map[string]interface{}, 6)
	familyShips["father"] = "tom"
	familyShips["mother"] = "mary"
	familyShips["son"] = "smith"

	p := Person{
		Name:   "jake",
		Age:    27,
		Height: 176,
		Skill:  &skills,
		FamilyShip: &familyShips,
	}

	// json序列化操作
	bytes, err := json.Marshal(&p)
	if err != nil {
		fmt.Println("json marshal failed! error: ", err)
	}
	fmt.Printf("struct json result: %v \n", string(bytes))

	// json反序列化操作
	jsonStr := string(bytes)
	var newPerson Person
	err = json.Unmarshal([]byte(jsonStr), &newPerson)
	if err != nil {
		fmt.Println("json unmarshal failed! error: ", err)
	}
	fmt.Printf("json反序列化结果：%v", newPerson)
}

func testMapJson()  {
	m := make(map[string]interface{})
	m["name"] = "tom"
	m["age"] = 68
	m["skills"] = []string{"足球", "钢琴", "吉他"}

	familyShips := make(map[string]interface{}, 6)
	familyShips["father"] = "tom"
	familyShips["mother"] = "mary"
	familyShips["son"] = "smith"

	m["familyShip"] = familyShips

	// json序列化操作
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal failed! error: ", err)
	}
	fmt.Printf("map json result: %v \n", string(bytes))

	// json反序列化操作
	jsonStr := string(bytes)
	var newMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &newMap)
	if err != nil {
		fmt.Println("json unmarshal failed! error: ", err)
	}
	fmt.Printf("json反序列化结果：%v", newMap)
}

func testSliceJson()  {
	var s []map[string]interface{}

	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 22
	m1["skills"] = []string{"篮球", "跳舞"}
	s = append(s, m1)


	familyShips := make(map[string]interface{}, 6)
	familyShips["father"] = "tom"
	familyShips["mother"] = "mary"
	familyShips["son"] = "smith"

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 34
	m2["skills"] = []string{"足球", "钢琴", "吉他"}
	m2["familyShip"] = familyShips
	s = append(s, m2)

	// json序列化操作
	bytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed! error: ", err)
	}
	fmt.Printf("slice json result: %v \n", string(bytes))

	// json反序列化操作
	jsonStr := string(bytes)
	var newSlice []map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &newSlice)
	if err != nil {
		fmt.Println("json unmarshal failed! error: ", err)
	}
	fmt.Printf("json反序列化结果：%v", newSlice)
}

func main() {
	testStructJson()
	fmt.Println()
	fmt.Println()

	testMapJson()
	fmt.Println()
	fmt.Println()

	testSliceJson()
}
