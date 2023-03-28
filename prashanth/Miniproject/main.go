package main

import (
	"fmt"
	"os"
	a "project/bulkload"
	b "project/marksprocess"
	c "project/webservice"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("studentmarks.csv")
	if err != nil {
		fmt.Print(err)
	}

	a.DataLoad(file)
	fmt.Println("data loaded successfully")

	b.GradesProcessor()
	fmt.Println("grades updated successfully")

	router := gin.Default()
	router.GET("/result/:id", c.GetResult)
	router.Run("localhost:8080")

}
