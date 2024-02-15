package models

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClientID     string    `gorm:"primaryKey" json:"ClientID"`
	FirstName    string    `json:"FirstName"`
	LastName     string    `json:"LastName"`
	Email        string    `json:"Email"`
	Phone        string    `json:"Phone"`
	Address      string    `json:"Address"`
	BusinessName string    `json:"BusinessName"`
	Status       string    `json:"Status"`
	ValidTill    time.Time `json:"ValidTill"`
}

func (c *Client) CreateUniqueID() {
	rand.Seed(time.Now().UnixNano())

	// You can customize this string based on your requirements
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a slice to store the characters of the random string
	result := make([]byte, 25)

	// Populate the slice with random characters
	for i := 0; i < 25; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}

	c.ClientID = string(result)
}
