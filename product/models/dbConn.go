package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Global instance DB connection
var DB *gorm.DB

// Initialize DB connection with PostgreSQL
func InitDB(cfg Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	// log.Println("Connected to Database.")

	// Migrates the User model by calling AutoMigrate.
	// This creates the necessary table in the database if it doesn't exist.
	if err := db.AutoMigrate(&Product{}); err != nil {
		panic(err)
	}

	// fmt.Println("Migrated database")

	DB = db
}
