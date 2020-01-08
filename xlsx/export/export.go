package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"time"
)

type User struct {
	Id int
	Name string
	Age int
	Sex int
}

func main() {
	t := time.Now()
	data := getDataFromDb()
	//fmt.Printf("data: %v\n", data)

	excel, err := generateExcel(data)
	since := time.Since(t)
	log.Printf("cost time: %v", since)
	checkErr(err)
	fmt.Printf("success: %v", excel)
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
	}

	for _, user := range data {
		contentRow := sheet.AddRow()
		var sex string
		if user.Sex == 0 {
			sex = "男"
		} else {
			sex = "女"
		}

		contentRow.AddCell().Value = strconv.Itoa(user.Id)
		contentRow.AddCell().Value = user.Name
		contentRow.AddCell().Value = strconv.Itoa(user.Age)
		contentRow.AddCell().Value = sex
	}

	err = file.Save("xlsx/export/user1.xlsx")
	if err != nil {
		return false, err
	}
	return true, nil
}

func getDataFromDb() []User {
	db, err := sql.Open("mysql", "root:13396095889@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("select * from user")
	checkErr(err)

	var data []User
	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
		logErr(err)

		data = append(data, user)
	}
	return data
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
func logErr(err error)  {
	if err != nil {
		log.Printf( "error log: %v \n", err)
	}
}
