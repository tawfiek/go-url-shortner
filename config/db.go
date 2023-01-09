package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"url-shortnner/models"
)

func InitDB(env Env) *gorm.DB {

	dsn := env.DBUser + ":" + env.DBPassword + "@tcp" + "(" + env.DBHost + ":" + env.DBPort + ")/" + env.DBName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Migrate(db)

	checkError(err)

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.URL{})

	checkError(err)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
