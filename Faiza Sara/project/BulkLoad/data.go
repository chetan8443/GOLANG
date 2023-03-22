package BulkLoad

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConn() *sql.DB {
	db, err := sql.Open("mysql", "root:Faiza@22@(127.0.0.1:3306)/student")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ReadData(file *os.File) {
	db := CreateConn()
	db.Exec("drop table students")
	db.Exec("create table students(Sid varchar(20), Sname varchar(20),Marks int)")

	csvFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvFile {

		_, err = db.Exec("insert into students values(?,?,?)", line[0], line[1], line[2])
		if err != nil {
			fmt.Println(err)
		}

	}
}
