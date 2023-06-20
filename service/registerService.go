package service

import (
	"socialNetwork/dto"
	"socialNetwork/model"
)

func RegisterUser(loginDto dto.LoginDto) bool {

	user := model.User{Email: loginDto.Email}

	AddUser(user)
	loginData := model.LoginData{Email: loginDto.Email, Password: loginDto.Password}
	model.DB.Create(&loginData)

	return true
}
