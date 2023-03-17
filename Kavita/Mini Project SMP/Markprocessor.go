	package main

	import (
		"database/sql"
		"fmt"
		"log"

		_ "github.com/mattn/go-sqlite3"
	)

	func main() {
		// Set up the database connection
		db, err := sql.Open("sqlite3", "../db/mydatabase.db")
		if err != nil {
			log.Fatal("DB error:", err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT studentid, studentname,studentmarks FROM students")
		if err != nil {
			log.Fatal("Error 1: ", err)
		}
		defer rows.Close()

		_, err = db.Exec("CREATE TABLE IF NOT EXISTS results_table (studentid INTEGER, studentname TEXT, result TEXT)")
		if err != nil {
			log.Fatal("Error 2: ", err)
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatal("Error 4: ", err)
		}

		for rows.Next() {
			var studentname string
			var studentmarks int
			var studentid string
			err = rows.Scan(&studentid, &studentname, &studentmarks)
			if err != nil {
				log.Fatal(err)
			}
			var result string
			if studentmarks >= 70 {
				result = "Pass with Distinction"
			} else if studentmarks > 40 {
				result = "Pass"
			} else {
				result = "Fail"
			}

			_, err = tx.Exec("INSERT INTO results_table (studentid, studentname, result) VALUES (?, ?, ?)", studentid, studentname, result)
			if err != nil {
				log.Fatal("Error 3: ", err)
			}
		}
		tx.Commit()

		fmt.Println("Results processed and inserted into results_table successfully.")
	}
