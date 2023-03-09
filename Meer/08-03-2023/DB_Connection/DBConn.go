// program to connect with mysql database 
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DBName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

var db *sql.DB

var config Config

func connect() {
	file1, err := os.Open("./dbconfig.json")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	defer file1.Close()

	decoder := json.NewDecoder(file1)
	config = Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to parse configuration file: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.User, config.Password, config.Host, config.Port)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database Connected Successfully !!!")

	db.Exec("CREATE DATABASE " + config.DBName)

	dsn = dsn + config.DBName
	db, err = sql.Open("mysql", dsn)

	file, err := ioutil.ReadFile("./sample.sql")

	queries := strings.Split(string(file), ";")

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func main() {
	connect()

}
