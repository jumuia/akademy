package repository

import "gorm.io/gorm"

func NewGORMUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return &gormUserRepository{
		DB: db,
	}
}
