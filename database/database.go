package database

import (
	"api-books/database/migrations"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=books password=admin"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("error:", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(10)
	config.SetConnMaxLifetime(10)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
