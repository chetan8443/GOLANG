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
	//fmt.Print("Reading the json file....")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error:%d", err)
	}

	user := viper.GetString("user")
	password := viper.GetString("password")
	host := viper.GetString("host")
	port := viper.GetString("port")
	dbname := viper.GetString("dbname")
	net := viper.GetString("net")
	addr := viper.GetString("addr")

	//opening the config file of json
	mysqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + "net" + net + "addr" + addr + " sslmode=disable"
	fmt.Print(mysqlInfo)
	db1, err := sql.Open("mysql", mysqlInfo)
	//**************************************
	if err != nil {
		fmt.Println("Check your config file...", err)
		return nil
	}
	//****************************************
	db1.Exec("DROP DATABASE studentinfo")
	_, err = db1.Exec("create database" + dbname)
	if err != nil {
		fmt.Print("Unable to create database", err)
	}
	configinfo := mysqlInfo + "database=" + dbname
	fmt.Println(configinfo)
	db2, err := sql.Open("mysql", configinfo)
	if err != nil {
		fmt.Println("Unable to read Json file", err)
	}
	defer db.Close()
	return db2
}

func MarksProcessor() {
	// fmt.Println("\nConnecting with database....")
	// db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/studentinfo")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// defer db.Close()
	// // fmt.Print("**********\n\n", arrays)
	// fmt.Println("Connected successfully...")
	// db.Exec("CREATE TABLE IF NOT EXISTS STUDENT (sid varchar(20) primary key , sname varchar(20),smarks int,result varchar(100))")
	// for _, info := range arrays {
	// 	//fmt.Println(info[0], info[1], info[2])
	// 	db.Exec("INSERT INTO STUDENT (sid,sname,smarks) values (info[0],info[1],info[2])")
	Connect()
}
