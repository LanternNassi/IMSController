package database

import (
	"github.com/LanternNassi/IMSController/internal/models"

	"context"
)

func (c Client) AddInstallation(ctx context.Context, installation *models.Installation) (*models.Installation, error) {

	result := c.DB.WithContext(ctx).Create(&installation)

	if result.Error != nil {
		return nil, result.Error
	}

	return installation, nil

}

func (c Client) UpdateInstallation(ctx context.Context, installation *models.Installation, id string) (*models.Installation, error) {
	result := c.DB.WithContext(ctx).Where(id).Updates(installation)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return installation, nil
}

func (c Client) GetInstallations(ctx context.Context, params *models.Installation) ([]models.Installation, error) {
	var installations []models.Installation

	result := c.DB.WithContext(ctx).Where(params).Find(&installations)

	if result.Error != nil {
		return nil, result.Error
	}

	return installations, nil
}

func (c Client) GetInstallationById(ctx context.Context, id string) (*models.Installation, error) {
	installation := &models.Installation{}

	result := c.DB.WithContext(ctx).First(installation, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return installation, nil
}

func (c Client) DeleteInstallation(ctx context.Context, id string) (bool, error) {

	installation, installation_error := c.GetInstallationById(ctx, id)

	if installation_error != nil {
		return false, installation_error
	}

	result := c.DB.WithContext(ctx).Delete(installation)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
