package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type User struct {
	Id int
	Name string
	Age int
	Sex int
}

func main() {
	var data []User
	var user User

	for i := 1; i < 10; i++ {
		user.Id = i
		user.Name = fmt.Sprintf("name%d", i)
		user.Age = i
		user.Sex = i % 2
		data = append(data, user)
	}

	ok, err := generateExcel(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}

func generateExcel(data []User) (bool, error) {
	titleRowInfo := make([]string, 0)
	titleRowInfo = append(titleRowInfo, "编号")
	titleRowInfo = append(titleRowInfo, "用户名")
	titleRowInfo = append(titleRowInfo, "年龄")
	titleRowInfo = append(titleRowInfo, "性别")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("user1")
	if err != nil {
		return false, err
	}

	titleRow := sheet.AddRow()
	for _, titleInfo := range titleRowInfo  {
		titleRow.AddCell().Value = titleInfo

		//titleRow.AddCell().SetFormat("")
	}

	//sexSlice := []string { "男", "女", "未知" }

	for _, user := range data {
		contentRow := sheet.AddRow()
		//var sex string
		//if user.Sex == 0 {
		//	sex = "男"
		//} else {
		//	sex = "女"
		//}

		//contentRow.AddCell().Value = strconv.Itoa(user.Id)
		//contentRow.AddCell().Value = user.Name
		//contentRow.AddCell().Value = strconv.Itoa(user.Age)
		//contentRow.AddCell().Value = sex

		//contentRow.WriteSlice(user, 1)
		contentRow.WriteStruct(&user, -1)
	}

	err = file.Save("xlsx/basic/user1.xlsx")
	if err != nil {
		return false, err
	}
	return true, nil
}