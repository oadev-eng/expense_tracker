package repository

import (
	"tracker/database"
	"tracker/models"
)


type UserRepository interface {
	GetUserByUserName(username string) (*models.User, error)
	CreateUser(user *models.User) error

}

type UserRepo struct{}

func (r *UserRepo) GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}