package database

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (c Client) GetBills(ctx context.Context, params *models.Bill) ([]models.Bill, error) {
	var bills []models.Bill

	result := c.DB.WithContext(ctx).Where(params).Order("ID DESC").Find(&bills)

	return bills, result.Error
}

func (c Client) GetBillsByDate(ctx context.Context, field string, comparator string, time_var time.Time, client_id string) ([]models.Bill, error) {
	var bills []models.Bill

	var result *gorm.DB

	result = c.DB.WithContext(ctx).Where(field+comparator+"?", time_var).Order("ID DESC").Find(&bills)

	if client_id != "" {
		result = c.DB.WithContext(ctx).Where(field+comparator+"?", time_var).Where("client_id = ?", client_id).Order("ID DESC").Find(&bills)
	}

	return bills, result.Error
}

func (c Client) AddBill(ctx context.Context, bill *models.Bill) (*models.Bill, error) {

	result := c.DB.WithContext(ctx).Create(&bill)

	if result.Error != nil {

		return nil, result.Error
	}

	return bill, nil
}

func (c Client) UpdateBill(ctx context.Context, bill *models.Bill, id string) (*models.Bill, error) {
	result := c.DB.WithContext(ctx).Where("ID = ?", id).Updates(models.Bill{
		BackupCount: bill.BackupCount,
		BackupSize:  bill.BackupSize,
		TotalCost:   bill.TotalCost,
		Billed:      bill.Billed,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
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
