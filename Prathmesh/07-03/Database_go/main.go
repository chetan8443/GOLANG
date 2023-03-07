package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("CRUD operation in golang with Postgres Sql")

	//database connection
	db, err := sql.Open("postgres", "postgres://postgres:Prathm@1387@localhost/prathmesh?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// creating table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, title TEXT, author TEXT, year INTEGER)")
	if err != nil {
		panic(err)
	}
	fmt.Println("table creation sucessfully")

	//Insert data:
	title := "The Great Gatsby"
	author := "F. Scott Fitzgerald"
	year := 1925
	id := 5
	_, err = db.Exec("INSERT INTO books (title, author, year) VALUES ($1, $2, $3)", title, author, year)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data inserted sucessfully")

	//Query data:
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var author string
		var year int

		err = rows.Scan(&id, &title, &author, &year)
		if err != nil {
			panic(err)
		}

	}
	fmt.Println(id, title, author, year)

}
