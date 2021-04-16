package database

import (
	"fmt"
	"ovh-availability/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// github.com/mattn/go-sqlite3
func ConnectDB() {
	var err error
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.ServerInfo{})
	fmt.Println("Database Migrated")
}
