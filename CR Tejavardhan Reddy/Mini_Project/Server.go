package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRequest() {
	r := gin.Default()
	r.GET("/", getAll)
	r.GET("/getByid/:id", getByID)
	r.RUN(":8081")
}

func getAll(c *gin.Context) {
	fmt.Println("HI")
	data, err := DB.Query("SELECT * FROM student ")

}

func getByID(c *gin.Context) {
	Id := r.URL.Query()["id"]
	fmt.Println(Id)
}
