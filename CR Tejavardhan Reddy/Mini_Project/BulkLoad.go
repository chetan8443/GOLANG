package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// type students struct {
// 	sid    string
// 	sname  string
// 	smarks int64
// }

// var s []students
var arrays [][]string

func OpenFile() {
	//opening csv file
	studentInfo, err := os.Open("./StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%T", studentInfo)
	reader := csv.NewReader(studentInfo)
	arrays, _ = reader.ReadAll()
	//fmt.Printf("%T", arrays)
	// for _, record := range arrays {
	// 	// a, _ := strconv.ParseInt(record[0], 10, 64) //converting into the int64 from string
	// 	b, _ := strconv.ParseInt(record[2], 10, 64)
	// 	s = append(s, students{sid: record[0], sname: record[1], smarks: b})
	// }
	// //fmt.Printf("%+v\n", s)
	fmt.Println(arrays)
}
