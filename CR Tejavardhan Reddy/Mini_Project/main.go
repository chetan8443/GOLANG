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
	router := gin.Default()
	router.GET("/marks/:id", w.GetDetails)
	router.Run(":4041")
	fmt.Println("Port established...")
}
