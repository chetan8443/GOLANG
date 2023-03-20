package bulkload

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	//connection to MYSQL database
	db, err := sql.Open("mysql", "root:Priyanka@gmail.com(127.0.0.1:3306)/priyanka")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func DataLoad(file *os.File) {
	db := Connection()
	//creating the student table
	db.Exec("drop table students")
	db.Exec("create table students(id varchar(255), name varchar(255),marks int)")

	// reading csv file line by line
	csvRec, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	//iterrating csv file

	for _, line := range csvRec {
		//query for inserting csv data to student table

		_, err = db.Exec("insert into students values(?,?,?)", line[0], line[1], line[2])
		if err != nil {
			fmt.Println(err)
		}

	}
}
