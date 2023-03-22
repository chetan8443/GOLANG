package MarkProcessing

import (
	"fmt"
	b "student/BulkLoad"
)

func CalcData() {

	db := b.CreateConn()
	db.Exec("drop table Results")
	db.Exec("create table Results(Sid varchar(20), Result varchar(50))")
	res, err := db.Query("select * from students")
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		var Sid, Sname string
		var marks int
		err = res.Scan(&Sid, &Sname, &marks)
		if err != nil {
			fmt.Println(err)
		}
		if marks >= 70 {
			res := "Pass with Distinction"
			db.Exec("insert into results values(?,?)", Sid, res)

		} else if marks > 40 {
			res := "Pass"
			db.Exec("insert into results values(?,?)", Sid, res)

		} else {
			res := "Fail"
			db.Exec("insert into results values(?,?)", Sid, res)
		}

	}
}
