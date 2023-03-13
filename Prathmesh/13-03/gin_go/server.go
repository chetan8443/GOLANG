package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/getAlbums", getAlbums)
	router.POST("/addData", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.PUT("/update/:id", updateAlbum)
	router.DELETE("/delete/:id", deleteAlbum)
	router.Run("localhost:8080")
}
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	//add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album added successfully"})

}
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if strconv.Itoa(a.ID) == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedalbum album
	if err := c.BindJSON(&updatedalbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, album := range albums {
		if strconv.Itoa(album.ID) == id {
			updatedalbum.ID = album.ID
			albums[i] = updatedalbum
			c.JSON(http.StatusOK, updatedalbum)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})

}
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, album := range albums {
		if strconv.Itoa(album.ID) == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})

}
