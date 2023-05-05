package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	MiddleName  string
	LastName    string
	DOB         *time.Time
	PhoneNumber string
	Email       string
}

type gormUserRepository struct {
	DB *gorm.DB
}

// Create implements UserRepository
func (r *gormUserRepository) Create(u *User) error {
	tx := r.DB.Create(u)
	if tx.Error != nil {
		log.Printf("unable to create user: %v", tx.Error)
		err := fmt.Errorf("user creation failed: %w", tx.Error)
		return err
	}
	log.Printf("Number of user records created is %d", tx.RowsAffected)
	return nil
}

// DeleteUser implements UserRepository
func (r *gormUserRepository) DeleteUser(id string) error {
	panic("unimplemented")
}

// FetchAllUsers implements UserRepository
func (r *gormUserRepository) FetchAllUsers() ([]*User, error) {
	panic("unimplemented")
}

// FindUser implements UserRepository
func (r *gormUserRepository) FindUser(id string) (*User, error) {
	user := User{}
	tx := r.DB.First(&user, id)
	if tx.Error != nil {
		log.Printf("unable to create user: %v", tx.Error)
		err := fmt.Errorf("user creation failed: %w", tx.Error)
		return nil, err
	}
	return &user, nil
}

// UpdateUser implements UserRepository
func (r *gormUserRepository) UpdateUser(*User) error {
	panic("unimplemented")
}
