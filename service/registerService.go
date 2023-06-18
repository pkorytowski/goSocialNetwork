package service

import "socialNetwork/model"

func RegisterUser(email string, password string) bool {

	user := model.User{Email: email}

	AddUser(user)
	loginData := model.LoginData{Email: email, Password: password}
	model.DB.Create(&loginData)

	return true
}
