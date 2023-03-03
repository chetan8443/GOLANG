// program to for create table , insert data, fetch data, and delete data from mysql
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:meer@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	_, err = db.Exec("Create table if not exists student (firstName varchar(20), lastName varchar(20), emailId varchar(20))")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table created Successfully ")
	}
	// insert, _ := db.Query("INSERT INTO `go`.`student` (`firstName`, `lastName`, `emailId`) VALUES ('Meer','Syed','fd@gmail.com');")
	// defer insert.Close()
	// fmt.Println("Data Inserted Successfully")

	rows, err := db.Query("Select * from student")

	for rows.Next() {
		var firstName, lastName, emailId string
		rows.Scan(&firstName, &lastName, &emailId)
		fmt.Println(firstName, " ", lastName, " ", emailId)

	}

	result, _ := db.Exec("delete from student where firstName = 'Meer'")
	fmt.Println(result.RowsAffected())
}
