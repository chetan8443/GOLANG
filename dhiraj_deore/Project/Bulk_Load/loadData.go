package Bulk_Load

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB { // making connectionn to database
	// provide all credentials to connect db server
	db, err := sql.Open("mysql", "root:Dhiraj@123@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db

}

func LoadData(file *os.File) {
	db := Connect() // connect to db
	// create a new table students
	db.Exec("drop table students")
	db.Exec("create table students(id varchar(255),name varchar(255),marks int)")

	//reading csv file line by line
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	//iterate over csv file line by line
	for _, line := range csvLines {
		//execute query to insert data into students table
		_, err = db.Exec("insert into students values(?,?,?)", line[0], line[1], line[2])
		if err != nil {
			fmt.Println(err)
		}
	}
}
