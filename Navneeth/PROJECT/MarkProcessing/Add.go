package MarkProcessing

import (
	"fmt"
	b "student/BulkLoad"
)

// type StudentMarks struct {
// 	Slno  string
// 	Name  string
// 	Marks string
// }

// var marks = []StudentMarks{}

func CalcData() {

	db := b.CreateConn()
	db.Exec("drop table Results")
	//db.Exec("create table Results(Slno varchar(20), Name varchar(20),Marks varchar(20), Result varchar(50))")
	db.Exec("create table Results(Slno varchar(20), Result varchar(50))")
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
		if marks >= 70 {
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
