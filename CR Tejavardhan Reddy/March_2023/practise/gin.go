package main

import "net/http"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HEllo world")
	})
}
