package Webserver

import (
	"database/sql"
	"fmt"
	"net/http"
	a "v1/MarkProcessor"

	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB
)

func GetDetails(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// id = string(id)
	var db = a.Connect()
	data, err := db.Query("SELECT result FROM result1 WHERE ID = ?", id)
	fmt.Println(data)
	for data.Next() {

		// result := ""
		var result string
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
