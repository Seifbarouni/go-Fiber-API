package data

import (
	"os"
	m "projects/Go-Fiber/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func InitializeDB() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = database.AutoMigrate(&m.User{}); err != nil {
		panic(err)
	}
	if err = database.AutoMigrate(&m.Article{}); err != nil {
		panic(err)
	}
	DB = database

}
