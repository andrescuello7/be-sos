package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnect(dns string) {
	var error error
	DB, error = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if error != nil {
		log.Fatal("Error connecting")
	} else {
		log.Println("Database connection successful")
	}
}
