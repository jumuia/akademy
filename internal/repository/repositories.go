package repository

import "gorm.io/gorm"

type UserRepository interface {
	Create(*User) error
	FindUser(id string) (*User, error)
	FetchAllUsers() ([]*User, error)
	UpdateUser(*User) error
	DeleteUser(id string) error
}

func NewGORMUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return &gormUserRepository{
		DB: db,
	}
}
