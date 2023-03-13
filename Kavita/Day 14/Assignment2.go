package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    
	Name string 
}

var users []User = []User{
	User{ID: 1, Name: "Alice"},
	User{ID: 2, Name: "Bob"},
	User{ID: 3, Name: "Charlie"},
}

func main(){
	router := gin.Default()

	router.PUT("/users/:id", updateUser)

	router.Run(":8080")
}

func updateUser(c *gin.Context) {
	fmt.Println("hello")// just for testing 
	id := c.Param("id")
	var updatedUser User

	err := c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
