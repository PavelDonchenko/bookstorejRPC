package config

import (
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dns := "pavel:mysqlpaha100688@tcp(127.0.0.1:3306)/testdb2?charset=utf8mb4&parseTime=True&loc=Local"
	d, err :=
		gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println(db)
}

func GetDB() *gorm.DB {
	return db
}
