package MarksProcessor

import (
	"database/sql"
	"fmt"
	"strconv"
	a "v1/BulkLoad"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)

func Connect() *sql.DB {
	//connection of database
	fmt.Println("Connecting with the database.....")
	viper.SetConfigName("configDB")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	fmt.Print("Reading the json file....\n\n")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error(unable to read Json file):", err)
	}

	user := viper.GetString("user")
	password := viper.GetString("password")
	dbname := viper.GetString("Dbname")
	net := viper.GetString("net")
	addr := viper.GetString("addr")
	var mysqlInfo string = user + ":" + password + "@" + net + "(" + addr + ")/"
	db1, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		fmt.Println("Fatal error(fail to check configfile):", err)
	}
	//creation of database
	db1.Query("DROP DATABASE studentinfo")
	_, err = db1.Query("create database " + dbname)
	if err != nil {
		fmt.Println("Fatal error(unable to create database):\n", err)
	}

	configinfo := mysqlInfo + dbname
	db2, err := sql.Open("mysql", configinfo)
	if err != nil {
		fmt.Println("Fatal error(unable to create database):", err)
	}
	return db2
}

func MarksProcessor() {
	//reading the file and building a table with the database
	DB = Connect()
	s := a.OpenFile()
	fmt.Println("Connected successfully...")
	DB.Query("CREATE TABLE STUDENT(sid varchar(200) primary key,sname varchar(200), marks int)")
	fmt.Println("Table created")
	fmt.Println("Inserting the data......")
	for i := 0; i < 5; i++ {
		DB.Query("INSERT INTO STUDENT(sid,sname,marks)values(?,?,?)", s[i][0], s[i][1], s[i][2])
		DB.Query("CREATE TABLE RESULT1(sid varchar(20)primary key,result varchar(200))")
		id := s[i][0]
		marks := s[i][2]
		x, _ := strconv.ParseInt(marks, 10, 64)
		//Manipulation of the Query
		if x >= 70 {
			DB.Query("INSERT INTO RESULT1(sid,result)values(?,'Pass With Distinction')", id)
		} else if x < 70 && x >= 40 {
			DB.Query("INSERT INTO RESULT1(sid,result)values(?,'Pass')", id)
		} else {
			DB.Query("INSERT INTO RESULT1(sid,result)values(?,'Fail')", id)
		}
	}
}
