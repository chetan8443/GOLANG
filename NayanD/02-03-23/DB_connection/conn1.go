//creating mysql database connectivity using golang and inserting data to it.

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open a database connection
	db, err := sql.Open("mysql", "root:nayan29@tcp(127.0.0.1:3306)/go_demo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Close the database connection after the function has finished executing
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Test the database connection

	fmt.Println("Successfully connected to the database.")

	insert, err := db.Query("INSERT INTO `go_demo`.`user` (`user_id`,`Name`, `email`) VALUES (2,'Arbind','Akumar@gmail.com');")
	defer insert.Close()
	fmt.Println("Data Inserted Successfully .")
}
