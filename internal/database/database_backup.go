package database

import (
	"github.com/LanternNassi/IMSController/internal/models"

	"time"

	"context"
)

func (c Client) Getbackups(ctx context.Context, params *models.Backup) ([]models.Backup, error) {
	var backups []models.Backup
	result := c.DB.WithContext(ctx).Where(params).Order("ID DESC").Find(&backups)
	for i := range backups {
		if err := c.DB.WithContext(ctx).Model(&backups[i]).Association("Bill").Find(&backups[i].Bill); err != nil {
			return nil, err
		}
	}
	return backups, result.Error
}

func (c Client) GetBackUpsByDate(ctx context.Context, field string, comparator string, time_var time.Time) ([]models.Backup, error) {
	var backups []models.Backup

	result := c.DB.WithContext(ctx).Where(field+comparator+"?", time_var).Order("ID DESC").Find(&backups)

	return backups, result.Error
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

func (c Client) DeleteBackUpById(ctx context.Context, id string) (bool, error) {
	result := c.DB.WithContext(ctx).Delete(&models.Backup{}, id).Unscoped()

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
