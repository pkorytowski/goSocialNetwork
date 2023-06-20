package model

import "time"

type Post struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	AuthorID  int       `json:"author_id" gorm:"foreignKey:AuthorID;references:ID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
