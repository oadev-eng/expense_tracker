package service

import (
	"errors"
	"tracker/middleware"
	"tracker/models"
	"tracker/repository"
	"tracker/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(user *models.User) error {
	// check if user exist already
	_, err := s.Repo.GetUserByUserName(user.Username)
	if err == nil {
		return errors.New("username already in use")
	}

	// hash the password
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashpass

	// call the create method
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginUser(request models.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByUserName(request.Username)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
