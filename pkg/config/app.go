package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	//Add .env file and setup MYSQL environment variables 	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/altschool", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"))
	
	d, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
