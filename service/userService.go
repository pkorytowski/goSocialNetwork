package service

import (
	"socialNetwork/model"
)

func AddUser(user model.User) model.User {
	model.DB.Create(&user)
	return user
}

func GetUsers() []model.User {
	var users []model.User
	model.DB.Find(&users)
	return users
}
