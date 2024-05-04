package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("POSTGRES_DSN"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&URL{})
	if err != nil {
		return
	}

	DbClient = database
}
