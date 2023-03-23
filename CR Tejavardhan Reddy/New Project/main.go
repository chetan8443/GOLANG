package main

import (
	"encoding/csv"
	"fmt"
	"os"

	A "new/BulkLoad"
	B "new/MarksProcessor"
	w "new/WebService"

	"github.com/gin-gonic/gin"
	// _ "github.com/gin-gonic/gin"
)

type Marks struct {
	ID          string
	StudentName string
	Marks       string
}

func main() {

	marks := []Marks{}

	var db = A.Connect()

	file, err := os.Open("BulkLoad/StudentMarks.csv")
	if err != nil {
		panic(err)
	}

	df := csv.NewReader(file)
	data, err := df.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		marks = append(marks, Marks{ID: value[0], StudentName: value[1], Marks: value[2]})
	}

	fmt.Printf("%+v", marks)
	for i := 0; i < len(marks); i++ {
		_, err = db.Exec("insert into result1(ID, StudentName, Marks) values(?, ?, ?)", marks[i].ID, marks[i].StudentName, marks[i].Marks)
		if err != nil {
			fmt.Println("unable to insert the data", err)
		}
	}
	fmt.Println("All Records Inserted...")

	B.MarksProcessor()
	fmt.Println("Result Inserted...")

	router := gin.Default()

	router.GET("/marks/:id", w.GetResultByID)
	router.Run("localhost:5001")
}
