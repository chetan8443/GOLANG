package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type student struct {
	sid    string
	sname  string
	smarks string
}

var s []student

func OpenFile() {
	//opening csv file
	studentInfo, err := os.Open("./StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%T", studentInfo)
	reader := csv.NewReader(studentInfo)
	arrays, _ := reader.ReadAll()
	fmt.Println(arrays)
	for _, record := range arrays {
		// a, _ := strconv.ParseInt(record[0], 10, 64) //converting into the int64 from string
		// b, _ := strconv.ParseInt(record[2], 10, 64)
		s = append(s, student{sid: record[0], sname: record[1], smarks: record[2]})
	}
	fmt.Printf("%+v\n", s)
}
