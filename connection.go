package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//db connection
var urlDSN = "root:root22YCA:@tcp(localhost:3306)/golang?parseTime=true"

func Connect() {
	connection, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&Users{})
}
