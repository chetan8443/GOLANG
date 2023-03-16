package WebService

import (
	"fmt"
	"net/http"
	b "student/BulkLoad"

	"github.com/gin-gonic/gin"
)

// func Api() {

// 	router := gin.Default()
// 	router.GET("/result/:id", ShowRes)
// 	router.Run("localhost:5000")
// }

func ShowRes(c *gin.Context) {
	fmt.Printf("show res called")
	id := c.Param("id")
	db := b.CreateConn()
	res, err := db.Query("select Result from Results where Slno =?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		var result string
		res.Scan(&result)
		response := map[string]string{"Result": result}
		c.IndentedJSON(http.StatusOK, response)
	}
}
