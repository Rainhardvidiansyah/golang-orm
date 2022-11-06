package service

import (
	"go-orm/models"
	"go-orm/repository"
)

type IUserService interface {
	SaveUser(models.User) (models.User, error)
}

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(Repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: Repository,
	}
}

func (userService *UserService) SaveUser(user *models.User) (models.User, error) {
	return *userService.Repository.CreateUser(user), nil
}
