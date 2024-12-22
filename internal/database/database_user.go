package database

import (
	"github.com/LanternNassi/IMSController/internal/models"

	"context"
)

func (c Client) AddUser(ctx context.Context, user *models.User) (*models.User, error) {

	result := c.DB.WithContext(ctx).Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (c Client) UpdateUser(ctx context.Context, user *models.User, id string) (*models.User, error) {
	result := c.DB.WithContext(ctx).Where(id).Updates(models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Verified: user.Verified,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (c Client) GetUserById(ctx context.Context, id string) (*models.User, error) {

	user := new(models.User)
	result := c.DB.WithContext(ctx).Where(id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (c Client) GetUsers(ctx context.Context, user *models.User) ([]models.User, error) {

	var users []models.User

	result := c.DB.WithContext(ctx).Find(&users, user)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (c Client) DeleteUser(ctx context.Context, id string) error {
	result := c.DB.WithContext(ctx).Where("ID = ?", id).Delete(&models.User{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
