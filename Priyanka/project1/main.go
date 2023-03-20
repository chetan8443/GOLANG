package main

import (
	"fmt"
	"os"
	b "project1/MarksProcessor"
	c "project1/WebService"
	a "project1/bulkload"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("Student.csv")
	if err != nil {
		fmt.Print(err)
	}
	a.DataLoad(file)
	fmt.Println("Data loaded successfully")

	b.GradesProcessor()
	fmt.Println("grades updated suceesfully")

	router := gin.Default()
	router.GET("/result/:id", c.GetResult)
	router.Run("localhost:8080")
}
