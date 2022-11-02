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
		return nil, errors.New("Tidak bisa save")
	} else {
		repository.DB.Create(&users)
	}
	return users, nil
}
