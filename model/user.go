package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email" gorm:"unique_index"`
	Password    string `json:"password"`
	Interests   string `json:"interests"`
	Hobby       string `json:"hobby"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}
