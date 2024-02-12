package database

import (
	"github.com/LanternNassi/IMSController/internal/models"

	"context"
)

func (c Client) Getbackups(ctx context.Context, params *models.Backup) ([]models.Backup, error) {
	var backups []models.Backup
	var trimmedbackups []models.Backup

	result := c.DB.WithContext(ctx).Where(params).Find(&backups)

	for _, backup := range backups {

		// Creating a new slice to avoid information overload during transfer
		backup.Backup = []byte{}

		trimmedbackups = append(trimmedbackups, backup)
	}

	return trimmedbackups, result.Error
}

func (c Client) AddBackup(ctx context.Context, backup *models.Backup) (*models.Backup, error) {

	result := c.DB.WithContext(ctx).Create(&backup)

	if result.Error != nil {

		return nil, result.Error
	}

	return backup, nil
}

func (c Client) GetBackUpById(ctx context.Context, id string) (*models.Backup, error) {
	backup := &models.Backup{}
	result := c.DB.WithContext(ctx).Where("ID = ?", id).First(&backup)

	if result.Error != nil {
		return nil, result.Error
	}

	return backup, nil

}
