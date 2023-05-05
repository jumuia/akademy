package service

import "github.com/jumuia/akademy/internal/repository"

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		Repo: repo,
	}
}
