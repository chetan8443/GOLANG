package webservice

import (
	"fmt"
	"net/http"
	a "project/bulkload"
	"github.com/gin-gonic/gin"
)

func GetResult(c *gin.Context) {
	id := c.Param("id")
	db := a.Connection()

	res, err := db.Query("select result from results where id =?", id)
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
