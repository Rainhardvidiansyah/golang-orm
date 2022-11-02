package repository

import (
	"errors"
	"go-orm/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(models.User) models.User
	FindUserByID(ID int) (models.User, error)
	SaveAllUser([]models.User) (*[]models.User, error)
	DeleteUserById(ID int) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (repository *UserRepository) CreateUser(user *models.User) *models.User {
	repository.DB.Create(&user)
	return user
}

func (repository *UserRepository) FindUserByID(ID int) *models.User {
	var user *models.User
	repository.DB.First(&user, ID)
	return user
}

func (repository *UserRepository) SaveAllUser(users *[]models.User) (*[]models.User, error) {
	if users == nil {
		return nil, errors.New("Cannot save users")
	} else {
		repository.DB.Create(&users)
	}
	return users, nil
}

func (repository *UserRepository) DeleteUserById(ID int) error {
	var user models.User
	repository.DB.Delete(&user, ID)
	return nil
}
