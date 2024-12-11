package models

import (
	"gorm.io/gorm"
)

type Backup struct {
	gorm.Model
	ClientID string `gorm:"primaryKey" json:"ClientID"`
	Name     string `json:"Name"`
	Backup   string `json:"Backup"`
	Size     int64  `json:"Size"`
	Bill     uint   `json:"Bill"`
}
