package WebService

import (
	"fmt"
	"net/http"
	b "new/BulkLoad"

	"github.com/gin-gonic/gin"
)

func GetResultByID(c *gin.Context) {
	id := c.Param("id")

	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	id = string(id)
	var db = b.GetMySQLDB()
	data, err := db.Query("select Result from Result1 where ID=?", id)
	fmt.Println(data)
	fmt.Println("idiot", data)
	for data.Next() {
		result := ""
		data.Scan(&result)
		response := map[string]string{"result": result}
		c.IndentedJSON(http.StatusOK, response)
	}
	data1, err1 := db.Query("select * from Result1 where ID=?", id)
	fmt.Println("wantes", *data1)
	b.CheckError(err)
	b.CheckError(err1)

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
