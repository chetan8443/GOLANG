package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArtist(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, artist)
}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
