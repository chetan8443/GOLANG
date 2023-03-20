package marksprocessor

import (
	"fmt"
	a "project/bulkload" 
	_ "github.com/go-sql-driver/mysql"
)

func FilterData() {

	db := a.Connection()
	db.Exec("drop table results")
	db.Exec("create table results(id varchar(255), result varchar(255))")

	res, err := db.Query("select * from students")
	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		var id, name string
		var marks int
		err = res.Scan(&id, &name, &marks)
		if err != nil {
			fmt.Println(err)
		}

		if marks > 65 {
			res := "First Class"
			db.Exec("insert into results values(?,?)", id, res)

		} else if marks > 40 {
			res := "Second Class"
			db.Exec("insert into results values(?,?)", id, res)

		} else {
			res := "Fail"
			db.Exec("insert into results values(?,?)", id, res)
		}

	}

}
