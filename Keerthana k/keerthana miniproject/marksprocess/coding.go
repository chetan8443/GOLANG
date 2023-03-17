package marksprocessor

import (
	"fmt"
	a "project/bulkload" // importing bulkload package for connection

	_ "github.com/go-sql-driver/mysql"
)

func GradesProcessor() {

	db := a.Connection()
	db.Exec("drop table results")
	db.Exec("create table results(id varchar(255), result varchar(255))")

	res, err := db.Query("select * from students")
	if err != nil {
		fmt.Println(err)
	}

	//iterating each record
	for res.Next() {
		var id, name string
		var marks int
		// scanning values of query into variables
		err = res.Scan(&id, &name, &marks)
		if err != nil {
			fmt.Println(err)
		}

		// assigning grades to student according to marks
		if marks > 69 {
			res := "Pass with Distinction"
			db.Exec("insert into results values(?,?)", id, res)

		} else if marks > 40 {
			res := "Pass with A grade"
			db.Exec("insert into results values(?,?)", id, res)

		} else {
			res := "You are Fail"
			db.Exec("insert into results values(?,?)", id, res)
		}

	}

}
