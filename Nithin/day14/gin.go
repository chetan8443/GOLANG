package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book = []Book{
	Book{ID: 1, Title: "1984", Author: "George Orwell"},
	Book{ID: 2, Title: "Brave New World", Author: "Aldous Huxley"},
	Book{ID: 3, Title: "Fahrenheit 451", Author: "Ray Bradbury"},
}

func main() {
	r := gin.Default()
	// Update a book
	r.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if strconv.Itoa(book.ID) == id {
				var updatedBook Book
				if err := c.ShouldBindJSON(&updatedBook); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "Invalid request body",
					})
					return
				}
				updatedBook.ID = book.ID
				books[i] = updatedBook
				c.JSON(http.StatusOK, gin.H{
					"message": "Book updated",
					"book":    updatedBook,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
	})

	// Delete a book
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if strconv.Itoa(book.ID) == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message": "Book deleted",
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
	})

	// Run the server
	r.Run("localhost:8080")
}
