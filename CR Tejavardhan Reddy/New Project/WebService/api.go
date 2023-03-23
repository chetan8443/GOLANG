package WebService

import (
	"fmt"
	"net/http"
	b "new/BulkLoad"

	"github.com/gin-gonic/gin"
)

func GetResultByID(c *gin.Context) {
	id := c.Param("id")
	id = string(id)
	var db = b.Connect()
	data, err := db.Query("select Result from Result1 where ID=?", id)
	fmt.Println(data)
	for data.Next() {
		result := ""
		data.Scan(&result)
		response := map[string]string{"result": result}
		c.IndentedJSON(http.StatusOK, response)
	}
	if err != nil {
		panic(err)
	}
}
