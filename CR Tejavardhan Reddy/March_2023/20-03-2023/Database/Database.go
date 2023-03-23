package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Connect() *sql.DB {
	//Establishing the dataBase connection for mysql
	db, err := sql.Open("mysql", "root:1234@@tcp(127.0.0.1:3306)/employee/result?parseTime=true")
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
