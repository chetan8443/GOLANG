package main

import (
	"database/sql"
	a "v/Database"
)

var DB *sql.DB

func main() {
	// 	r := gin.New()
	// 	// r := gin.Default()

	// 	r.GET("/", func(c *gin.Context) {
	// 		c.String(200, "HEllo")
	// 	})
	// 	r.Run(":8093")
	// }

	// func GET(c *gin.Context) {
	Database()

}
func Database() *sql.DB {

	a.Connect()
	DB.Query("CREATE DATABASE employee")
	DB.Query("CREATE TABLE Employee(eid int primary key,ename varchar(200))")
	for i := 1; i <= 5; i++ {
		DB.Exec("INSERT INTO Employee(eid,ename)values(?,?)", i, string(i+65))
	}
	return DB
}
