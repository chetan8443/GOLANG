package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Sid    int
	Sname  string
	Marks  int
	Result string
}

func main() {

	// Connect
	db, err := sql.Open("mysql", "root:Faiza@22@(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Retrieve
	rows, err := db.Query("SELECT Sid, Sname, Marks FROM studentinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process
	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Sid, &student.Sname, &student.Marks)
		if err != nil {
			log.Fatal(err)
		}
		if student.Marks >= 70 {
			student.Result = "Pass with Distinction"
		} else if student.Marks >= 40 {
			student.Result = "Pass"
		} else {
			student.Result = "Fail"
		}
		students = append(students, student)
	}

	// Add
	_, err = db.Exec("ALTER TABLE studentinfo ADD COLUMN result VARCHAR(50)")
	if err != nil {
		log.Fatal(err)
	}

	// Update
	for _, student := range students {
		_, err := db.Exec("UPDATE studentinfo SET result = ? WHERE sid = ?", student.Result, student.Sid)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done")
}
