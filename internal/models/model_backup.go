package models

import (
	"gorm.io/gorm"
)

type Backup struct {
	gorm.Model
	ClientID string `gorm:"primaryKey" json:"ClientID"`
	Name     string `json:"Name"`
	Backup   []byte `json:"Backup"`
	Size     int64  `json:"Size"`
	Billed   bool   `json:"Billed"`
}
