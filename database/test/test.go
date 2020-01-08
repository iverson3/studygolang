package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, _ := sql.Open("mysql", "root:13396095889@tcp(localhost:3306)/test?charset=utf8")

	//queryOneRow(db)
	transactionOperation(db)
}

func queryOneRow(db *sql.DB) {
	var name string
	var age int
	err := db.QueryRow("select name,age from user where id=?", 33).Scan(&name, &age)
	//drivers := sql.Drivers()
	//columnType := sql.ColumnType{}

	switch {
	case err == sql.ErrNoRows:
		log.Println("no user")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("name: %s; age: %d", name, age)
	}
}

func transactionOperation(db *sql.DB) {
	tx, err := db.Begin()
	checkErr(err)

	var id int
	name := "name22"
	age := 22
	rows, err := tx.Query("select id from user where name=? and age=?", name, age)
	errRollback(err, tx)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		errRollback(err, tx)
	}
	fmt.Println(id)

	res, err := tx.Exec("update user set sex=? where id=?", 1, id)
	errRollback(err, tx)
	affect, err := res.RowsAffected()
	errRollback(err, tx)

	fmt.Printf("affect row: %d\n", affect)

	if affect == 1 {
		fmt.Println("success, commit")
		err = tx.Commit()
		if err != nil {
			err = tx.Rollback()
			checkErr(err)
		}
	} else {
		fmt.Println("fail, rollback")
		err = tx.Rollback()
		checkErr(err)
	}
}

func errRollback(err error, tx *sql.Tx)  {
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("rollback failed: %v", err)
		}
		return
	}
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}