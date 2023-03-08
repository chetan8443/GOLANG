package main

import ( // importing libraries
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct { // creating a struct data type
	Id        string `json:"id"`
	Item      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Item: "clean room", Completed: false},
	{Id: "2", Item: "washing", Completed: false},
	{Id: "3", Item: "walking", Completed: false},
}

func getTodos(c *gin.Context) { // converting data into json format
	c.IndentedJSON(http.StatusOK, todos)
}

func main() { // using GET to call API
	router := gin.Default()
	router.GET("/todos", getTodos)

	router.Run("localhost:8080")
}
