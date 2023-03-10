package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:dhiraj@tcp(localhost:3306)/test?parseTime=true"
var err error

func DataHandler() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&Employee{})

}
