package main

import (
	"fmt"
	"os"
	r "student/BulkLoad"
	p "student/MarkProcessing"
	a "student/WebService"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		fmt.Print(err)
	}
	r.ReadData(file)
	fmt.Println("File is Loaded")
	p.CalcData()
	fmt.Println("Results are calculated")

	router := gin.Default()
	router.GET("/result/:id", a.ShowRes)
	router.Run("localhost:8082")
}
