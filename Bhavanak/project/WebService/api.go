package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMySQLDB() *sql.DB {
	//connection
	db, err := sql.Open("mysql", "root:9704348918@Bc@(127.0.0.1:3306)/student?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func GetResult(c *gin.Context) {
	id := c.Param("id")
	db := getMySQLDB()

	res, err := db.Query("select result from results where id =?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		var result string
		res.Scan(&result)
		response := map[string]string{"Result": result}
		c.IndentedJSON(http.StatusOK, response)
	}
}
func main() {
	router := gin.Default()
	router.GET("/result/:id", GetResult)
	router.Run("localhost:8080")

	router.Run() //listen and serve on 0.0.0.0:8080
}
