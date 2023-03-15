package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Bload() {
	file, _ := os.Open("StudentMarks.csv") //Open opens the named file for reading
	df := csv.NewReader(file)              //NewReader returns a new Reader that reads from r.
	slices, _ := df.ReadAll()              //ReadAll reads all the remaining records from r. Each record is a slice of fields
	fmt.Println(slices)

	for _, sl := range slices {
		id, _ := strconv.Atoi(sl[0])
		mk, _ := strconv.Atoi(sl[2])
		Studs = append(Studs, Student{Id: id, Name: sl[1], Marks: mk}) //Append built-in function appends elements to the end of a slice

	}
	fmt.Println(Studs)
}
