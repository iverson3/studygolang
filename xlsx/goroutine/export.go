package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Id int
	Name string
	Age int
	Sex int
}

// 分页大小： 每10w条数据作为一个sheet
const sheetRows int = 100000

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

	var wg sync.WaitGroup

	var sheetNums = len(data) / sheetRows
	//sheetNums = 1
	for i := 1; i <= sheetNums; i++ {
		wg.Add(1)
		go createSheet(&wg, titleRowInfo, data[(i - 1) * sheetRows:i * sheetRows], i)
	}

	wg.Wait()
	return true, nil
}

func createSheet(wg *sync.WaitGroup, titleRowInfo []string, data []User, i int)  {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet(fmt.Sprintf("user%d", i))

	titleRow := sheet.AddRow()
	for _, titleInfo := range titleRowInfo  {
		titleRow.AddCell().Value = titleInfo
	}

	for _, user := range data {
		contentRow := sheet.AddRow()
		//fmt.Println(i)
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
	log.Printf("before save: %d", i)
	if true {
		file.Save(fmt.Sprintf("xlsx/goroutine/user%d.xlsx", i))
	}
	log.Printf("after save: %d", i)

	wg.Done()
}

func getDataFromDb() []User {
	db, err := sql.Open("mysql", "root:13396095889@tcp(47.107.149.234:3306)/test?charset=utf8")
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
