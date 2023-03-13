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

	// Get all books
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	})

	// Get a specific book by ID
	r.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, book := range books {
			if strconv.Itoa(book.ID) == id {
				c.JSON(http.StatusOK, gin.H{
					"book": book,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
	})

	// Add a new book
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
			return
		}
		book.ID = len(books) + 1
		books = append(books, book)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Book added",
			"book":    book,
		})
	})

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
