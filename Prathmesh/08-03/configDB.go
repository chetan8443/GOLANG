package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)


func Connect() *sql.DB {

	fmt.Println("Connecting to database ... ")
	viper.SetConfigName("configdb")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	fmt.Println("Reading configuration from json file ... ")
	err := viper.ReadInConfig() 
	if err != nil {
		panic(fmt.Errorf("Fatal error reading json config file: %w \n", err))
	}

	user := viper.GetString("user") // viper is to read both from configuration file and environment variables.
	password := viper.GetString("password")
	host := viper.GetString("host")
	port := viper.GetString("port")
	dbname := viper.GetString("dbname")
	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " sslmode=disable"
	fmt.Printf(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Check your config file, Database not connected ...")
		//fmt.Printf(err)
		return nil
	}

	_, err = db.Exec("create database " + dbname)
	conninfob := psqlInfo + " database=" + dbname
	db2, err := sql.Open("postgres", conninfob)
	file, err := ioutil.ReadFile("./dbschema.sql")

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		result, err := db2.Exec(request)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
	defer db.Close()

	DB = db2
	return DB
}
func main() {
	Connect()
}
