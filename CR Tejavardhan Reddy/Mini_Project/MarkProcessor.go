package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	db *sql.DB
)

func Connect() *sql.DB {
	fmt.Println("Connecting with the database.....")
	viper.SetConfigName("configDB")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	fmt.Print("Reading the json file....\n\n")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error(unable to read Json file):", err))
	}

	user := viper.GetString("user")
	password := viper.GetString("password")
	dbname := viper.GetString("Dbname")
	net := viper.GetString("net")
	addr := viper.GetString("addr")
	// fmt.Println("\n*******************************************")
	//opening the config file of json
	var mysqlInfo string = user + ":" + password + "@" + net + "(" + addr + ")/"
	// fmt.Print(mysqlInfo)
	db1, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(fmt.Errorf("Fatal error(fail to check configfile):", err))
	}
	db1.Exec("DROP DATABASE studentinfo")
	_, err = db1.Exec("create database " + dbname)
	if err != nil {
		panic(fmt.Errorf("Fatal error(unable to create database):\n", err))
	}

	configinfo := mysqlInfo + dbname
	//fmt.Println(configinfo)
	db2, err := sql.Open("mysql", configinfo)
	if err != nil {
		panic(fmt.Errorf("Fatal error(unable to create database):", err))
	}
	defer db1.Close()
	return db2
}

func MarksProcessor() {
	db = Connect()
	fmt.Println("Connected successfully...")
	db.Exec("CREATE TABLE STUDENT(sid varchar(200) primary key,sname varchar(200), marks int,result varchar(200))")
	fmt.Println("Table created")
	fmt.Println("Inserting the data......")
	for i, data := range arrays {
		db.Exec("INSERT INTO STUDENT(sid,sname,marks)values(?,?,?)", data[0], data[1], data[2])
		fmt.Println(i+1, "th row inserted")
	}
	fmt.Println("The data inserted..")
}
