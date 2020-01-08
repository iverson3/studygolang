package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("mysql database")

	db, err := connectDb()
	checkErr(err)
	defer func() {
		// 当发生panic时，所在goroutine的所有defer会被执行
		// 所以下面代码的db 可能会是nil; 如果是nil调用Close()会出错
		err = db.Close()
		checkErr(err)
	}()

	//simpleTransaction(db)

	// insert
	doInsert(db)
	// update
	//doUpdate(db)
	// select
	//doSelect(db)
	// delete
	//doSelect(db)
}

func connectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:13396095889@tcp(localhost:3306)/test?charset=utf8")
	return db, err
}

func simpleTransaction(db *sql.DB)  {
	tx, err := db.Begin()
	checkErr(err)

	var id int
	name := ""
	age := 22
	rows, err := tx.Query("select id from user where name=? and age=?", name, age)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		checkErr(err)
	}
	err = tx.QueryRow("select name, age from user where id = ?", 3).Scan(&name, &age)
	checkErr(err)

	res, err := tx.Exec("update user set sex=? where id=?", 1, id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)

	if affect == 1 {
		err = tx.Commit()
		checkErr(err)
	} else {
		err = tx.Rollback()
		checkErr(err)
	}
}

func doInsert(db *sql.DB) {
	stmt, err := db.Prepare("insert user set name=?, age=?, sex=?")
	checkErr(err)
	for i := 0; i < 700000; i++ {
		res, err := stmt.Exec(fmt.Sprintf("name%d", i + 1), i + 1, i%2)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Printf("insert id: %d; affect rows: %d", id, affect)
	}
}

func doUpdate(db *sql.DB) {
	prepare, err := db.Prepare("update user set name=? where id=?")
	checkErr(err)
	res, err := prepare.Exec("new name", 2)
	checkErr(err)

	affectRow, err := res.RowsAffected()
	fmt.Printf("affect row: %d\n", affectRow)
}

func doSelect(db *sql.DB) {
	rows, err := db.Query("select * from user")
	checkErr(err)
	defer rows.Close()

	//fmt.Printf("%v", rows)
	for rows.Next() {
		var id, age, sex int
		var name string

		err := rows.Scan(&id, &name, &age, &sex)
		checkErr(err)

		fmt.Printf("user %d: %s %d %d", id, name, age, sex)
		fmt.Println()
	}
}

func doDelete(db *sql.DB) {
	stmt, err := db.Prepare("delete from user where id=?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(33)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("affect rows: %d\n", affect)
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
