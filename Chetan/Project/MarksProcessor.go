package main

import (
	"database/sql"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Marks  int
	Result string
}

var Db2 *sql.DB

var Studs []Student

func MarksProcessor() {

	Db2 = Connet()	//Package sql provides a generic interface around SQL (or SQL-like) databases.

	//Creating student table
	Db2.Exec("CREATE TABLE students (id int PRIMARY KEY,name varchar(50) NOT NULL,marks int, result varchar(50))")

	for i := 0; i < len(Studs); i++ {
		fmt.Println(Studs[i].Marks)

			//Base on marks inserting records in table
				if Studs[i].Marks >= 70 {
					fmt.Printf("Student %s Pass with Distinction\n", Studs[i].Name)
					Db2.Exec("INSERT INTO students (id, name, marks,result)VALUES (?, ?, ?,'Pass with Distinction')", Studs[i].Id, Studs[i].Name, Studs[i].Marks)

				} else if Studs[i].Marks < 70 && Studs[i].Marks >= 40 {
					fmt.Printf("Student %s Pass\n", Studs[i].Name)
					Db2.Exec("INSERT INTO students (id, name, marks,result)VALUES (?, ?, ?,'Pass')", Studs[i].Id, Studs[i].Name, Studs[i].Marks)

				} else {
					fmt.Printf("Student %s Fail\n", Studs[i].Name)
					Db2.Exec("INSERT INTO students (id, name, marks,result)VALUES (?, ?, ?,'Fail')", Studs[i].Id, Studs[i].Name, Studs[i].Marks)

				}
	}

	fmt.Println("Succefully Inserted !")

}
