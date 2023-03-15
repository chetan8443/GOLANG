package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// var (
// 	db *sql.DB		//Package sql provides a generic interface around SQL (or SQL-like) databases.
// )

func Connet() *sql.DB {

	//setting config file information with viper
	viper.SetConfigName("dbConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	fmt.Println("Reading config from json file...")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad json file\n", err))
	}

	//Assigning values from DB config JSON file
	user := viper.GetString("User")
	password := viper.GetString("Password")
	dbname := viper.GetString("DBName")
	Net := viper.GetString("Net")
	Addr := viper.GetString("Addr")

	//Makes connection string (dataSourceName) so we can connect to Sql
	var ConnectionString string = user + ":" + password + "@" + Net + "(" + Addr + ")/"

	//  Open opens a database specified by its database driver name and a driver-specific data source name
	db1, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		panic(fmt.Errorf("Fatal error1 to read json file\n", err))
	}

	db1.Exec("DROP DATABASE MARCH")                //Droppping database if exists
	_, err = db1.Exec("create database " + dbname) //Creating Database by given  name
	if err != nil {
		panic(fmt.Errorf("Fatal error2 to read json file\n", err))
	}

	//connectionString (DSN) for  opening specific dataBase
	conninfob := ConnectionString + dbname
	db2, err := sql.Open("mysql", conninfob)
	if err != nil {
		panic(fmt.Errorf("Fatal error to rad json file\n", err))
	}

	return db2

}
