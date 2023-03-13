package main

import ( 
	"net/http"
	"github.com/gin-gonic/gin"
)


type Book struct {
	ID int
	Title string

}

var books = []Book{
	{ID: 1, Title: "Bhagvat geeta"},
	{ID: 2, Title: "Ramayana"},
}


func getBooks (c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}



func main()  {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
	
}

        


