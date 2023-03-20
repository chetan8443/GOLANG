package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	sid   string
	sname string
	marks string
}

func getMySQLDB() *sql.DB {
	//connection
	db, err := sql.Open("mysql", "root:9704348918@Bc@(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	var db = getMySQLDB()
	studentinfo := []student{}
	//opening csv file
	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}
	//Reading csv file
	df := csv.NewReader(file)
	data, err := df.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	//iterating through the csv file
	for _, value := range data {
		studentinfo := append(studentinfo, student{sid: value[0], sname: value[1], marks: value[2]})
		//inserting csv data to studentinfo table

		_, err = db.Exec("insert into studentinfo values(?,?,?)", value[0], value[1], value[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", studentinfo)
	}

}
