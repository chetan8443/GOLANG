package BulkLoad

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type students struct {
	sid    string
	sname  string
	smarks int64
}

var s []students
var arrays [][]string

func OpenFile() [][]string {
	//opening csv file
	studentInfo, err := os.Open("BulkLoad/StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(studentInfo)
	arrays, _ = reader.ReadAll()
	for _, record := range arrays {
		b, _ := strconv.ParseInt(record[2], 10, 64)
		s = append(s, students{sid: record[0], sname: record[1], smarks: b})
	}
	//fmt.Printf("%+v\n", s)
	//fmt.Println(arrays)
	fmt.Println("The CSV data is loaded")
	return arrays
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		fmt.Println("Unable to connect the database")
	}
	return db
}
