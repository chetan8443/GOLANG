package BulkLoad

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConn() *sql.DB {
	db, err := sql.Open("mysql", "root:twentyfour@/test?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ReadData(file *os.File) {
	db := CreateConn()
	db.Exec("drop table students")
	db.Exec("create table students(Slno varchar(20), Name varchar(20),Marks int)")

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
