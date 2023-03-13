package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json : "id"`
	Item      string `json : "item"`
	Completed bool   `json : "completed"`
}

var todos = []todo{
	{Id: "1", Item: "golang assignment", Completed: false},
	{Id: "2", Item: "git push", Completed: false},
	{Id: "3", Item: "update work status", Completed: false},
}

// function to add new todos
func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

//funtion to show the todo list
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

//function to show the todos by id
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todo", getTodos)    //to get todo
	router.GET("/todo/:id", GetTodo) //to get todo by id
	router.POST("/todo", addTodos)   //to add new todo
	router.Run("localhost:5000")

}
