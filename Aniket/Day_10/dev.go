package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function to get all cources inforamation in JSON format as response
func getCourses(c *gin.Context) {
	// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
	c.IndentedJSON(http.StatusOK, courses)
}

// Function to add new cources inforamation in JSON format as post request
func addCourse(c *gin.Context) {
	var newCourse *Course

	// Call BindJSON to bind the received JSON to newAlbum.
	err := c.BindJSON(&newCourse)
	checkNilError(err)
	courses := append(courses, *newCourse)
	c.IndentedJSON(http.StatusCreated, courses)
}

// Function to check error in given function
func checkNilError(err error) {
	// If error is not NULL then panic function will terminate the current operation
	if err != nil {
		panic(err)
	}
}
