package main

import (
	"encoding/csv"
	"fmt"
	"os"

	b "new/BulkLoad"
	m "new/MarksProcessor"
	w "new/WebService"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Marks struct {
	ID          string
	StudentName string
	Marks       string
}

func main() {

	marks := []Marks{}

	var db = b.GetMySQLDB()

	file, err := os.Open("StudentMarks.csv")
	b.CheckError(err)

	df := csv.NewReader(file)
	data, err := df.ReadAll()
	b.CheckError(err)

	for _, value := range data {
		marks = append(marks, Marks{ID: value[0], StudentName: value[1], Marks: value[2]})
	}

	fmt.Println(marks)

	for i := 0; i < len(marks); i++ {
		_, err = db.Exec("insert into result(ID, StudentName, Marks) values(?, ?, ?)", marks[i].ID, marks[i].StudentName, marks[i].Marks)
		b.CheckError(err)
	}
	fmt.Println("All Records Inserted...")

	m.IngestRes()
	fmt.Println("Result Inserted...")

	router := gin.Default()

	router.GET("/marks/:id", w.GetResultByID)

	router.Run("localhost:5000")
}
