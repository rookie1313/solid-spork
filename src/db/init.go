package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"solid-spork/src/model"
)

func InitDB() *gorm.DB {
	config := gorm.Config{TranslateError: true, Logger: logger.Default}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=root password=root dbname=solid_spork port=5432 sslmode=disable TimeZone=UTC",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &config)
	if err != nil {
		log.Fatal("failed to init db")
	}

	// migrate db models
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to migrate table")
	}

	return db
}
