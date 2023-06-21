package model

type Follow struct {
	ID         int `json:"id" gorm:"primary_key auto_increment"`
	FollowsID  int `json:"follows_id" gorm:"foreignKey:FollowsID;references:ID"`
	FollowedID int `json:"followed_id" gorm:"foreignKey:FollowedID;references:ID"`
}
