package database

import (
	"github.com/LanternNassi/IMSController/internal/models"

	"context"
)

func (c Client) AddClient(ctx context.Context, client *models.Client) (*models.Client, error) {

	//Creating a unique id
	client.CreateUniqueID()

	result := c.DB.WithContext(ctx).Create(&client)

	if result.Error != nil {

		return nil, result.Error
	}

	return client, nil
}

func (c Client) UpdateClient(ctx context.Context, client *models.Client, id string) (*models.Client, error) {

	result := c.DB.WithContext(ctx).Where("ID = ?", id).Updates(models.Client{
		FirstName:    client.FirstName,
		LastName:     client.LastName,
		Email:        client.Email,
		Phone:        client.Phone,
		Address:      client.Address,
		BusinessName: client.BusinessName,
		Status:       client.Status,
	})

	if result.Error != nil {

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return client, nil

}

func (c Client) GetClientById(ctx context.Context, id string) (*models.Client, error) {
	client := &models.Client{}
	result := c.DB.WithContext(ctx).Where("ID = ?", id).First(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return client, nil

}

func (c Client) GetClients(ctx context.Context, params *models.Client) ([]models.Client, error) {
	var clients []models.Client
	result := c.DB.WithContext(ctx).Where(params).Find(&clients)

	if result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}
