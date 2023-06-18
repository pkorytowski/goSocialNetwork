package service

import (
	"golang.org/x/crypto/bcrypt"
	"socialNetwork/model"
	"socialNetwork/token"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginUser(email string, password string) (string, error) {

	var err error

	l := model.LoginData{}

	err = model.DB.Where("email = ?", email).First(&l).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, l.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	u := model.User{}
	err = model.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(uint(u.ID))

	if err != nil {
		return "", err
	}

	return "Bearer " + token, nil

}
