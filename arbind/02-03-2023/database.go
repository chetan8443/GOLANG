package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	insert, err := db.Query("INSERT INTO `godb`.`student` (`firstName`, `lastName`, `emailId`) VALUES ('arbind','chauhan','arbindchauhan00@gmail.com');")
	defer insert.Close()
	fmt.Println("Data Inserted Successfully")
}
