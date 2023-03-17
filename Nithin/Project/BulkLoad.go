package BulkLoad

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Nithin@16@(127.0.0.1:3306)/result")
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
