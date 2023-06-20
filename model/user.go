package model

type User struct {
	ID          int    `json:"id" gorm:"primary_key auto_increment"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email" gorm:"primary_key unique"`
	Interests   string `json:"interests"`
	Hobby       string `json:"hobby"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}
