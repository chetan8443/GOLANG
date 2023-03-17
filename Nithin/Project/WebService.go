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
	var db = b.GetMySQLDB()
	data, err := db.Query("select result from result1 where ID=?", id)
	fmt.Println(data)
	for data.Next() {
		result := ""
		data.Scan(&result)
		responce := map[string]string{"result": result}
		c.IndentedJSON(http.StatusOK, responce)
	}

	b.CheckError(err)

}
