// Write a program to use post method using gin
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string `json:"name"`
	City string `json:"city"`
}

var list []student

func main() {
	router := gin.Default()
	router.POST("/info", info)
	router.Run("localhost:8080")
}

func info(c *gin.Context) {
	var new student
	if err := c.BindJSON(&new); err != nil {
		return
	}
	list = append(list, new)
	fmt.Print(list)
	c.IndentedJSON(http.StatusOK, map[string]string{"messege": "Successfully added !!"})
}
