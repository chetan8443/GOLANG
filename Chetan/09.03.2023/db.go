package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	db *sql.DB
)

func Connet() *sql.DB {
	viper.SetConfigName("dbConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	fmt.Println("Reading config from json file...")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad json file\n", err))
	}
	user := viper.GetString("User")
	// host := viper.GetString("Host")
	password := viper.GetString("Password")
	// port := viper.GetString("Port")
	dbname := viper.GetString("DBName")
	Net := viper.GetString("Net")
	Addr := viper.GetString("Addr")

	var ConnectionString string
	ConnectionString=user+":"+password+"@"+Net+"("+Addr+")/"
	// mysqlinfo:=
	db1, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad1 json file\n", err))
	}
	_, err = db1.Exec("create database " + dbname)
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad2 json file\n", err))
	}

	conninfob := ConnectionString + dbname
	db2, err := sql.Open("mysql", conninfob)
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad json file\n", err))
	}

	dbyte, err := ioutil.ReadFile("./test.sql")
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad json file\n", err))
	}
	queries := strings.Split(string(dbyte), ";")

	for _, query := range queries {
		fmt.Println(query)
		_, err := db2.Exec(query)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println("result : ", result)

	}
	db = db2
	return db

}

func main() {
	Connet()
}
