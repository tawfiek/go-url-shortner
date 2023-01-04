package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dbConnection *gorm.DB

func InitDB(env Env) *gorm.DB {

	dsn := env.DBUser + ":" + env.DBPassword + "@tcp" + "(" + env.DBHost + ":" + env.DBPort + ")/" + env.DBName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	checkError(err)

	return db
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
