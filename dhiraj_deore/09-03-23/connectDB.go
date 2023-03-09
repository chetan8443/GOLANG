package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connectDB()
}

func connectDB() *sql.DB {
	type postgres struct {
		Host     string `json:"Host"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Port     string `json:"Port"`
		Dbname   string `json:"Dbname"`
	}
	//var conn map[string]string
	var conn postgres
	file, _ := os.Open("configDB.json")
	a, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(a, &conn)
	//fmt.Println(conn)
	//psqlinfo := "host=" + conn["Host"] + " port=" + conn["Port"] + " user=" + conn["User"] + " password=" + conn["Password"] + " dbname=" + conn["Dbname"] + " sslmode=disable"
	psqlinfo := "host=" + conn.Host + " port=" + conn.Port + " user=" + conn.User + " password=" + conn.Password + " dbname=" + conn.Dbname + " sslmode=disable"
	//fmt.Println(psqlinfo)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	fmt.Println("connected")
	result, _ := db.Query("select * from marks")
	for result.Next() {
		var a, b, c, d string
		result.Scan(&a, &b, &c, &d)
		fmt.Println(a, b, c, d)
	}
	return db
}
