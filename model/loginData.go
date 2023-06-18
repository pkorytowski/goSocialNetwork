package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginData struct {
	ID       int    `json:"id" gorm:"primary_key auto_increment"`
	Email    string `json:"email" gorm:"foreignKey:Email;references:Email"`
	Password string `json:"password"`
}

func (u *LoginData) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
