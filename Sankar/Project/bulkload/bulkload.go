package bulkload
import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
)
func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:sankar@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func LoadingData(file *os.File) {
	db := Connection()
	db.Exec("drop table students")
	db.Exec("create table students(id varchar(255), name varchar(255),marks int)")

	csvrecords, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, column := range csvrecords {
		_, err = db.Exec("insert into students values(?,?,?)", column[0], column[1], column[2])
		if err != nil {
			fmt.Println(err)
		}
	}
}
