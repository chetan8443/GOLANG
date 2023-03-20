package main

import (
	"fmt"
	"os"
	"net/http"
	a "project/bulkload"
	b "project/marksfilter"
	c "project/webservice"
	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("studentmarks.csv")
	if err != nil {
		fmt.Print(err)
	}

	a.LoadingData(file)
	fmt.Println("data loaded successfully")

	b.FilterData()
	fmt.Println("grades updated successfully")

	router := gin.Default()
	router.GET("/result/:id", c.GetResult)
	router.GET("/result", msg)
	router.GET("/details/:name", c.GetData)
	router.Run("localhost:8087")

}
func msg(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"messege": "Details of the Students"})
}
