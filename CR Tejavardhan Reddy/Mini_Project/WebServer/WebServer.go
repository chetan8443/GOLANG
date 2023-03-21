package Webserver

import (
	"fmt"
	"net/http"
	a "v1/BulkLoad"

	"github.com/gin-gonic/gin"
)

func GetDetails(c *gin.Context) {
	id := c.Param("id")
	//fmt.Println(id)
	id = string(id)
	var db = a.Connect()
	data, err := db.Query("SELECT * FROM result WHERE ID=?", id)
	fmt.Println(data)
	for data.Next() {
		result := ""
		data.Scan(&result)
		output := map[string]string{"result": result}
		c.IndentedJSON(http.StatusOK, output)
	}
	if err != nil {
		fmt.Println(err)
	}
}
