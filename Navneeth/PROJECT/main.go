package main

import (
	"fmt"
	"os"
	b "student/BulkLoad"
	m "student/MarkProcessing"
	w "student/WebService"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		fmt.Print(err)
	}
	b.ReadData(file)
	fmt.Println("File is Loaded")
	m.CalcData()
	fmt.Println("Results are calculated")

	router := gin.Default()
	router.GET("/result/:id", w.ShowRes)
	router.Run("localhost:5050")
}
