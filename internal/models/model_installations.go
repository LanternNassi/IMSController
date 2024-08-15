package models

import (
	"gorm.io/gorm"
)

type Installation struct {
	gorm.Model
	ClientID          string `json:"ClientID"`
	Installation_type string `json:"Installation_type"`
	Computer_name     string `json:"Computer_name"`
	IMS_version       string `json:"IMS_version"`
	Operating_system  string `json:"Operating_system"`
	RAM               string `json:"RAM"`
	Processor         string `json:"Processor"`
	Active            string `json:"Active"`
}
