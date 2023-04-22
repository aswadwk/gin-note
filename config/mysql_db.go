package config

import (
	"aswadwk/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file config")
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")

	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306" // Default port if not specified
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to database")

	db.AutoMigrate(&models.Activity{}, &models.Todo{})
	log.Println("Migration")

	return db
}