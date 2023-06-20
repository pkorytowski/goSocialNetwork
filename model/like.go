package model

type Like struct {
	ID     int `json:"id" gorm:"primary_key auto_increment"`
	UserID int `json:"user_id" gorm:"foreignKey:users references:id"`
	PostID int `json:"post_id" gorm:"foreignKey:posts references:id"`
}
