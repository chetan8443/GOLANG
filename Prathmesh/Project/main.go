package main

import (
	"fmt"
	"os"
	b "project/MarksProcessor"
	c "project/WebService"
	a "project/bulkload"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("StudentMarks.csv")
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
