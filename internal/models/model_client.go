package models

import (
	"math/rand"
	"time"
)

type Client struct {
	ClientID      string         `gorm:"primaryKey" json:"ClientID"`
	FirstName     string         `json:"FirstName"`
	LastName      string         `json:"LastName"`
	Email         string         `json:"Email"`
	Phone         string         `json:"Phone"`
	Address       string         `json:"Address"`
	BusinessName  string         `json:"BusinessName"`
	Status        string         `json:"Status"`
	ValidTill     time.Time      `json:"ValidTill"`
	Installations []Installation `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Installations"`
	Bills         []Bill         `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Bills"`
	Backups       []Backup       `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Backups"`
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
