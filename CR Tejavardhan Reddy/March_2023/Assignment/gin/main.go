package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the Gin router
	router := gin.Default()

	// Update a user
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User updated"})
	})

	// Delete a user
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted"})
	})

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
