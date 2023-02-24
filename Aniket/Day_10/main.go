package main

import "github.com/gin-gonic/gin"

type Course struct {
	CourseName  string  `json:"coursename"`
	CourseId    int     `json:"courseid"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"name"`
	Website    string `json:"website"`
}

var courses = []Course{
	{CourseId: 001, CourseName: "Golang", CoursePrice: 5000, Author: &Author{AuthorName: "Aniket", Website: "aniket.com"}},
	{CourseId: 002, CourseName: "Python", CoursePrice: 5040, Author: &Author{AuthorName: "Tushar", Website: "tushar.com"}},
}

func main() {
	router := gin.Default()

	router.GET("/get", getCourses)
	router.POST("/post", addCourse)

	router.Run("localhost:5000")
}
