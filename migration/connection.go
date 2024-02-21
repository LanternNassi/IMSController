package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB - Get gorm DB instance.
func NewDB() (*gorm.DB, error) {

	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading environment variables file")
	}

	dbport, err_conv := strconv.Atoi(os.Getenv("DBPORT"))

	if err_conv != nil {
		fmt.Println("Error converting string to int:", err_conv)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), dbport, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
