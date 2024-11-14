package db

import (
	"fmt"
	"log"

	"github.com/javimartzs/client.io/config"
	"github.com/javimartzs/client.io/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig) *gorm.DB {

	uri := fmt.Sprintf(
		`user=%s pass=%s dbname=%s host=%s port=%s`,
		config.DBUser, config.DBPass, config.DBName, config.DBHost, config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
