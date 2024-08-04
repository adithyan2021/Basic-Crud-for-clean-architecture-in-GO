package user

import "gorm.io/gorm"

type UserRegister struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Phone    string `json:"phone" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
