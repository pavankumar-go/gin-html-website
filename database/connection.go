package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	db, err := gorm.Open(sqlite.Open("gorm.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open sql connection...\n ", err.Error())
	}
	log.Println("connected!")
	DB = db
}

func GetDBConnection() *gorm.DB {
	if DB == nil {
		log.Println("Reinitializing DB connection...")
		initDB()
		return DB
	}
	return DB
}
