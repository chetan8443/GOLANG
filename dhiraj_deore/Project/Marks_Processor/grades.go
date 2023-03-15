package Marks_Processor

import (
	"fmt"
	b "project/Bulk_Load" //import Bulk Load package for database connection
)

func ProcessGrades() {
	db := b.Connect() // connect to database
	//create a new table for results
	db.Exec("drop table results")
	db.Exec("create table results(id varchar(255),result varchar(255))")

	res, err := db.Query("select * from students") // getting records from students table
	if err != nil {
		fmt.Println(err)
	}
	// iterate over each record
	for res.Next() {
		var id, name string
		var marks int
		err = res.Scan(&id, &name, &marks) // scan values of query into variables
		if err != nil {
			fmt.Println(err)
		}
		// assign grades according to marks
		if marks > 69 {
			res := "Pass with Distinction"
			db.Exec("insert into results values(?,?)", id, res)
		} else if marks > 40 {
			res := "Pass"
			db.Exec("insert into results values(?,?)", id, res)
		} else {
			res := "Fail"
			db.Exec("insert into results values(?,?)", id, res)
		}
	}
}
