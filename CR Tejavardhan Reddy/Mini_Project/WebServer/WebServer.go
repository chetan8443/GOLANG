package webserver

import (
	"fmt"
	"net/http"
	a "v1/MarkProcessor"

	"github.com/gin-gonic/gin"
)

func GetDetails(c *gin.Context) {
	id := c.Param("id")
	//fmt.Println(id)
	id = string(id)
	var db = a.Connect()
	data, err := db.Query("SELECT * FROM RESULT WHERE ID=?", id)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data)
	for data.Next() {
		var id int
		var result string

		data.Scan(&id, &result)
		if result != "" {
			c.IndentedJSON(http.StatusOK, gin.H{"id": id, "Result": result})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
}
