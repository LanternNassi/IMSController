package database

import (
	"context"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (c Client) GetBills(ctx context.Context, params *models.Bill) ([]models.Bill, error) {
	var bills []models.Bill

	result := c.DB.WithContext(ctx).Where(params).Find(&bills)

	return bills, result.Error
}

func (c Client) AddBill(ctx context.Context, bill *models.Bill) (*models.Bill, error) {

	result := c.DB.WithContext(ctx).Create(&bill)

	if result.Error != nil {

		return nil, result.Error
	}

	return bill, nil
}

func (c Client) GetBillById(ctx context.Context, id string) (*models.Bill, error) {
	bill := &models.Bill{}
	result := c.DB.WithContext(ctx).Where("ID = ?", id).First(&bill)

	if result.Error != nil {
		return nil, result.Error
	}

	return bill, nil
}
