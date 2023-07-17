package model

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := getDB()
	if err != nil {
		log.Fatalf("gorm open database error: %v", err)
	}
	DB = db
	DB.AutoMigrate(&Relay{}, &RelayEvent{})
}

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("POSTGRES_URL"),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if os.Getenv("POSTGRES_DEBUG") == "true" {
		db = db.Debug()
	}

	return db, nil
}
