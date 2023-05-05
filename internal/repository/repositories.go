package repository

import "gorm.io/gorm"

type UserRepository interface {
}

func NewGORMUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{
		DB: db,
	}
}

type gormUserRepository struct {
	DB *gorm.DB
}
