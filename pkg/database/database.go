package database

import (
	"fmt"
	"log"

	"github.com/ahmadalaik/be-invitation/config"
	"github.com/ahmadalaik/be-invitation/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.LoadConfig().DBHost,
		config.LoadConfig().DBUser,
		config.LoadConfig().DBName,
		config.LoadConfig().DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Template{},
		&models.Invitation{},
		&models.StoryImage{},
		&models.Guest{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	return db, err
}
