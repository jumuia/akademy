package service

import "github.com/jumuia/akademy/internal/repository"

type UserService interface {
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		Repo: repo,
	}
}

type userService struct {
	Repo repository.UserRepository
}
