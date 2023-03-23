package Webserver

import (
	"fmt"
	"net/http"
	a "v1/MarkProcessor"

	"github.com/gin-gonic/gin"
)

func GetDetails(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// id = string(id)
	var db = a.Connect()
	data, err := db.Query("SELECT result FROM result1 WHERE ID = ?", id)
	// fmt.Println(data)
	for data.Next() {

		result := ""
		data.Scan(&result)
		// fmt.Println("Vardh")
		output := map[string]string{"result": result}
		c.IndentedJSON(http.StatusOK, output)
	}
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Final")
}
