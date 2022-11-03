package service

import "go-orm/repository"

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(Repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: Repository,
	}
}
