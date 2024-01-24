package models

import (
	"gorm.io/gorm"
)

type Backup struct {
	gorm.Model
	ClientID string `gorm:"primaryKey" json:"ClientID"`
	Name     string `json:"Name"`
	Backup   string `json:"Backup"`
}
