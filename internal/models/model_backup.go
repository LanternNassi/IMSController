package models

import (
	"gorm.io/gorm"
)

type Backup struct {
	gorm.Model
	ClientID string `json:"ClientID"`
	Name     string `json:"Name"`
	Backup   string `json:"Backup"`
	Size     int64  `json:"Size"`
	BillID   uint   `gorm:"index" json:"BillID"`
	Bill     Bill   `gorm:"foreignKey:BillID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Bill"`
}
