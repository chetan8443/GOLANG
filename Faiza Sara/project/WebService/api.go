package WebService

import (
	"fmt"
	"net/http"
	r "student/BulkLoad"

	"github.com/gin-gonic/gin"
)

func ShowRes(c *gin.Context) {
	fmt.Printf("show res called")
	id := c.Param("id")
	db := r.CreateConn()
	res, err := db.Query("select Result from Results where Sid =?", id)
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
