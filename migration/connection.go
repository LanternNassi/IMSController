package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// NewDB - Get gorm DB instance.
func NewDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123456789 dbname=imscontroller port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}