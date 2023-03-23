package BulkLoad

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(localhost:3306)/studentinfo?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
