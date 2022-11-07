package service

import (
	"go-orm/models"
	"go-orm/repository"
)

type IUserService interface {
	SaveUser(user models.User) (models.User, error)
	FindUserByID(ID int) (models.User, error)
	SaveAllUser(user []models.User) (*[]models.User, error)
	DeleteUserById(ID int) error
	UpdateUserName(ID int, user *models.User) (*models.User, error)
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

func (userService *UserService) FindUserByID(ID int) *models.User {
	user := userService.Repository.FindUserByID(ID)
	return user
}

func (userService *UserService) SaveAllUser(users *[]models.User) (*[]models.User, error) {
	user, _ := userService.Repository.SaveAllUser(users)
	return user, nil
}

func (userService *UserService) DeleteUserById(ID int) error {
	err := userService.Repository.DeleteUserById(ID)
	return err
}

func (userService *UserService) UpdateUserName(ID int, user *models.User) (*models.User, error) {
	userName, _ := userService.Repository.UpdateUserName(ID, user)
	return userName, nil
}
