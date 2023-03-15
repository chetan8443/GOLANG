// API to show result of given id
package Web_Service

import (
	"fmt"
	"net/http"
	b "project/Bulk_Load" // importing Bulk_Load package to connect with database

	"github.com/gin-gonic/gin" //import gin
)

func GetResult(c *gin.Context) {
	id := c.Param("id") // taking parameter from url request
	db := b.Connect()   // connect to database
	// get result of given id from database
	res, err := db.Query("select result from results where id =?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() { // iterate over query result
		var result string
		res.Scan(&result) // scanning query result into result variable
		c.IndentedJSON(http.StatusOK, gin.H{"Result": result}) // sending response in json format
	}

}
