// Create a program with a GET API that retrieves data by Id.
package main

import (
	"io/ioutil"
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
		{Name: "Shreya", Id: 1, SurName: "patil"},
		{Name: "swara", Id: 2, SurName: "patil"},
		{Name: "Jayanta", Id: 3, SurName: "patil"}}
	r := gin.Default()

	r.GET("/getStudent/:id", getStudentById)
	r.GET("/getStudent", getStudents)
	r.DELETE("/delete/:id", dStudentById)
	r.PUT("/update", UpdateStudent)

	r.Run(":8080")

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

func dStudentById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for i, s := range student {
		if id == s.Id {
			student = append(student[:i], student[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "student details have been Removed successfully"})

			return
		}
	}

}

func UpdateStudent(c *gin.Context) {

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
