package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getMySQLDB() *sql.DB {
	//connection
	db, err := sql.Open("mysql", "root:9704348918@Bc@(127.0.0.1:3306)/student?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {

	var db = getMySQLDB()
	db.Exec("drop table results")
	db.Exec("create table results(sid varchar(255), result varchar(255))")

	res, err := db.Query("select * from studentinfo")
	if err != nil {
		fmt.Println(err)
	}

	//iterating each record
	for res.Next() {
		var sid, sname string
		var marks int
		// scanning values of query into variables
		err = res.Scan(&sid, &sname, &marks)
		if err != nil {
			fmt.Println(err)
		}

		// assigning grades to student according to marks
		if marks >= 70 {
			res := "Pass with Distinction"
			db.Exec("insert into results values(?,?)", sid, res)

		} else if marks >= 40 {
			res := "Pass with A grade"
			db.Exec("insert into results values(?,?)", sid, res)

		} else {
			res := "You are Fail"
			db.Exec("insert into results values(?,?)", sid, res)
		}

		//}
		fmt.Println("Table Inserted")

		// 	result, err := db.Query("select studentinfo.sid,studentinfo.sname,studentinfo.marks,results.result from (studentinfo) join (results) on (studentinfo.sid=results.sid)")
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	db.Exec("insert into results values(?,?,?,?)", sid, sname, marks, result)

		// }
	}
}
