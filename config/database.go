package config

import (
	"os"

	"github.com/ArwahDevops/gin-postgresql-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbURL := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
