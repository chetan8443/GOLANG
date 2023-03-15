package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Api() {

	r := gin.Default()
	r.GET("/getStudent", getStudents)
	r.GET("/getStudent/:id", getStudentById)
	r.Run(":8080") //Run attaches the router to a http.Server and starts listening and serving HTTP requests
}

func getStudents(c *gin.Context) {

	rows, _ := Db2.Query("SELECT * FROM students ") //Query executes a query that returns rows

	for rows.Next() {
		var name, result string
		var id, marks int
		if err := rows.Scan(&id, &name, &marks, &result); err != nil { //Scan copies the columns in the current row into the values pointed at variables
			fmt.Errorf("Got an error %v", err)
		}
		c.IndentedJSON(http.StatusOK, gin.H{"Student ID:": id, "Name": name, "result": result, "Marks :": marks})
	}

}

func getStudentById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	rows, _ := Db2.Query("SELECT marks,result FROM students where id = ?", id)

	for rows.Next() {
		var result string
		var marks int

		rows.Scan(&marks, &result)

		if result != "" {
			c.IndentedJSON(http.StatusOK, gin.H{"Result": result, "Marks": marks})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}
