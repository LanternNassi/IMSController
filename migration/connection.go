package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// NewDB - Get gorm DB instance.
func NewDB() (*gorm.DB, error) {
	//dsn := "host=localhost user=postgres password=123456789 dbname=imscontroller port=5432 sslmode=disable TimeZone=UTC"
	dsn := "host=102.134.147.233 user=tsaonokrtitbzytk password=4HLeh9LwP0aN+Dach0E1jt5I7K%N%5m7 dbname=edendsmwqscdihljjantklpo port=32761 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}