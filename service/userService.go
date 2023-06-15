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

func GetUserById(id int) model.User {
	var user model.User
	model.DB.First(&user, id)
	return user
}

func UpdateUser(user model.User) model.User {
	model.DB.Save(&user)
	return user
}

func DeleteUser(user model.User) {
	model.DB.Delete(&user)
}

func GetUserByEmail(email string) model.User {
	var user model.User
	model.DB.Where("email = ?", email).First(&user)
	return user
}
