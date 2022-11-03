package repository

import (
	"errors"
	"go-orm/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(models.User) models.User
	FindUserByID(ID int) (models.User, error)
	SaveAllUser(user []models.User) (*[]models.User, error)
	DeleteUserById(ID int) error
	UpdateUserName(ID int, user *models.User) (*models.User, error)
	UpdateUser(ID int, user *models.User) (*models.User, error)
	FindAllUser() (*[]models.User, error)
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

//This function only updates single column.
func (repository *UserRepository) UpdateUserName(ID int, user *models.User) (*models.User, error) {
	repository.DB.Model(user).Where("ID =?", ID).Update("UserName", user.UserName)
	return user, nil
}

//This function will updates multiple columns.
func (repository *UserRepository) UpdateUser(ID int, u *models.User) (*models.User, error) {
	err := repository.DB.Find(&u, ID)
	if err != nil {
		return nil, errors.New("No Data Found")
	}
	repository.DB.Model(&u).Where("ID=?", ID).Updates(models.User{UserName: u.UserName, Email: u.Email})
	return u, nil
}

func (respository *UserRepository) FindAllUser() (*[]models.User, error) {
	var user *[]models.User
	respository.DB.Find(&user)
	return user, nil
}
