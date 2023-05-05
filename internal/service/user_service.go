package service

import (
	"fmt"
	"log"
	"time"

	"github.com/jumuia/akademy/internal/repository"
)

type User struct {
	FirstName   string
	MiddleName  string
	LastName    string
	DOB         string
	PhoneNumber string
	Email       string
}

type userService struct {
	Repo repository.UserRepository
}

// Create implements UserService
func (s *userService) Create(u *User) error {
	dob, err := time.Parse("2006-01-02", u.DOB)
	if err != nil {
		log.Printf("unable to parse date of birth: %v", err)
		err = fmt.Errorf("unable to parse the date of birth: %w", err)
		return err
	}
	user := repository.User{
		FirstName:   u.FirstName,
		MiddleName:  u.MiddleName,
		LastName:    u.LastName,
		DOB:         &dob,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
	return s.Repo.Create(&user)
}

// DeleteUser implements UserService
func (s *userService) DeleteUser(id string) error {
	panic("unimplemented")
}

// FetchAllUsers implements UserService
func (s *userService) FetchAllUsers() ([]*User, error) {
	panic("unimplemented")
}

// FindUser implements UserService
func (s *userService) FindUser(id string) (*User, error) {
	u, err := s.Repo.FindUser(id)
	if err != nil {
		log.Printf("unable to find user: %v", err)
		return nil, err
	}
	user := User{
		FirstName:   u.FirstName,
		MiddleName:  u.MiddleName,
		LastName:    u.LastName,
		DOB:         u.DOB.String(),
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
	return &user, nil
}

// UpdateUser implements UserService
func (s *userService) UpdateUser(*User) error {
	panic("unimplemented")
}
