package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Sid   string
	Sname string
	Marks string
}

func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:Faiza@22@(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	var db = connection()
	// Open
	studentinfo := []Student{}
	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Read
	rd := csv.NewReader(file)
	data, err := rd.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range data {

		studentinfo := append(studentinfo, Student{Sid: value[0], Sname: value[1], Marks: value[2]})

		for i := 0; i < len(studentinfo); i++ {
			Sid, _ := strconv.Atoi(studentinfo[i].Sid)
			Marks, _ := strconv.Atoi(studentinfo[i].Marks)

			_, err := db.Exec("insert into studentinfo values(?,?,?)", Sid, studentinfo[i].Sname, Marks)
			if err != nil {
				log.Fatal(err)
			}

		}
		fmt.Printf("%+v\n", studentinfo)
	}
}
