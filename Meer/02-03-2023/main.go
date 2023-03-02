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

	insert, _ := db.Query("INSERT INTO `go`.`student` (`firstName`, `lastName`, `emailId`) VALUES ('Meer','Syed','meer@gmail.com');")
	defer insert.Close()
	fmt.Println("Data Inserted Successfully")
}
