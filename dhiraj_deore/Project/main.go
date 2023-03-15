package main

import (
	"fmt"
	"os"
	b "project/Bulk_Load"
	m "project/Marks_Processor"
	r "project/Web_Service"

	"github.com/gin-gonic/gin"
)

func main() {
	// open csv file
	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		fmt.Print(err)
	}

	// loading data of csv file into database
	b.LoadData(file)
	fmt.Println("data loaded successfully !!")

	// Adding results table and calculate result of each student according to marks
	m.ProcessGrades()
	fmt.Println("Grades updated successfully !!")

	// using gin package creating a api to show result of given id
	router := gin.Default()
	router.GET("/result/:id", r.GetResult)
	router.Run("localhost:8080")

}
