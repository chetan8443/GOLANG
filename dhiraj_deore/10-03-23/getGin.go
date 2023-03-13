// Write a program to write a msg using gin
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get_msg", msg)
	router.Run("localhost:8080")
}

func msg(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"messege": "Successfully running !!"})
}
