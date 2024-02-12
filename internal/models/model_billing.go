package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ClientID    string          `json:"ClientID"`
	BackupCount int             `json:"BackupCount"`
	TotalCost   decimal.Decimal `gorm:"type:numeric" json:"Size"`
}
