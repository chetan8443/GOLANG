package main

import (
	"fmt"

	a "v1/BulkLoad"
	b "v1/MarkProcessor"
	w "v1/WebServer"

	"github.com/gin-gonic/gin"
)

func main() {

	//assigning the port
	a.OpenFile()
	b.MarksProcessor()
	router := gin.Default()

	// fmt.Println("OKOK")
	router.GET("/marks/:id", w.GetDetails) //.Methods("GET")
	router.Run("127.0.0.1:5001")
	// fmt.Fprintln(w, "Hello")
	fmt.Println("Port established...")
}
