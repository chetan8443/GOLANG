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

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:dhiraj@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Dummy struct {
	sid    string
	name   string
	marks  string
	result string
}

func main() {
	var db = getMySQLDB()

	// courses := []model.StudentDetails{}
	courses := []Dummy{}

	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var de string
	var res string

	for _, value := range data {

		// fmt.Println(v)

		// fmt.Printf("var1 = %T\n", value[i])
		de = value[2]
		val, _ := strconv.Atoi(de)

		if val > 75 {
			res = "passed with distinction"
		} else if val < 75 && val > 60 {
			res = "passed with grade A"
		} else if val < 60 && val > 50 {
			res = "passed with grade B"
		} else if val < 50 && val > 40 {
			res = "passed with grade C"
		} else {
			res = "failed"
		}
		courses = append(courses, Dummy{sid: value[0], name: value[1], marks: value[2], result: res})

	}
	for j := 1; j < len(courses); j++ {
		sid, _ := strconv.Atoi(courses[j].sid)
		marks, _ := strconv.Atoi(courses[j].marks)

		_, err := db.Exec("insert into studentdetail(id, name, marks,result) values(?,?,?,?)", sid, courses[j].name, marks, courses[j].result)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("all data inserted")

}
