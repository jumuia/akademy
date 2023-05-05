package service

import "github.com/jumuia/akademy/internal/repository"

type UserService interface {
	Create(*User) error
	FindUser(id string) (*User, error)
	FetchAllUsers() ([]*User, error)
	UpdateUser(*User) error
	DeleteUser(id string) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		Repo: repo,
	}
}
