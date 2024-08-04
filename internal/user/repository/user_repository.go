package repository

import (
	user "github.com/shakezidin/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *user.UserRegister) error
	FindUserByUserName(username string) (*user.UserRegister, error)
}

type UserDataBaseInteraction struct {
	DB *gorm.DB
}

func (u *UserDataBaseInteraction) CreateUser(user *user.UserRegister) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserDataBaseInteraction) FindUserByUserName(username string) (*user.UserRegister, error) {
	var user *user.UserRegister
	if err := u.DB.Where("userrname = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDataBaseInteraction{
		DB: db,
	}
}
