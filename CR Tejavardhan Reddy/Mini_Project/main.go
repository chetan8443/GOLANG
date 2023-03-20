package main

import (
	"fmt"
	a "v1/BulkLoad"
	b "v1/MarkProcessor"
	w "v1/WebServer"

	"github.com/gin-gonic/gin"
)

func main() {
	a.OpenFile()
	b.MarksProcessor()
	//assigning the port
	router := gin.New()
	router.GET("/marks/:id", w.GetDetails)
	router.Run(":8081")
	fmt.Println("Port established...")
}
