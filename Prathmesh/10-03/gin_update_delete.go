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

var books []Book

func main() {
	r := gin.Default()

	// Create a new book
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book.ID = len(books) + 1
		books = append(books, book)

		c.JSON(http.StatusCreated, book)
	})

	// Get all books
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// Get a single book
	r.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, book := range books {
			if strconv.Itoa(book.ID) == id {
				c.JSON(http.StatusOK, book)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	})

	// Update a book
	r.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedBook Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, book := range books {
			if strconv.Itoa(book.ID) == id {
				updatedBook.ID = book.ID
				books[i] = updatedBook
				c.JSON(http.StatusOK, updatedBook)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	})

	// Delete a book
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if strconv.Itoa(book.ID) == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	})

	r.Run(":8080")
}
