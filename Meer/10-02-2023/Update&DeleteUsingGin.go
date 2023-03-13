// program to update and delete data by Id using Gin
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name    string `json:"Name"`
	Id      int    `json:"Id"`
	SurName string `json:"SurName"`
}

var student []Student

func main() {
	student = []Student{
		{Name: "Meer", Id: 1, SurName: "Syed"},
		{Name: "Fahad", Id: 2, SurName: "Khan"},
		{Name: "Ali", Id: 3, SurName: "Patel"}}
	r := gin.Default()

	r.GET("/student/:id", getStudentById)
	r.GET("/student", getStudents)
	r.DELETE("/student/:id", deleteStudentById)
	r.PUT("/student/:id", UpdateStudentById)

	r.Run(":8086")

}

func getStudents(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, student)
}

// func getStudent(w http.ResponseWriter, r *http.Request){

// }

func getStudentById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for i, s := range student {
		if id == s.Id {
			c.IndentedJSON(http.StatusOK, student[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func deleteStudentById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for i, s := range student {
		if id == s.Id {
			student = append(student[:i], student[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})

			return
		}
	}

}

func UpdateStudentById(c *gin.Context) {

	var std Student

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBind(&std)

	c.JSON(http.StatusOK, std)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, s := range student {
		if id == s.Id {
			student = append(student[:i], student[i+1:]...)
			student = append(student, std)

			c.JSON(http.StatusOK, gin.H{"message": "Student details have been Updated successfully"})

			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student details not Found"})

}
