package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ClientID    string          `json:"ClientID"`
	BackupCount int             `json:"BackupCount"`
	BackupSize  int64           `json:"BackupSize"`
	TotalCost   decimal.Decimal `gorm:"type:numeric" json:"TotalCost"`
	Billed      bool            `json:"Billed"`
}
