package model

import "time"

type Comment struct {
	ID              int       `json:"id" gorm:"primary_key auto_increment"`
	AuthorID        int       `json:"author_id" gorm:"foreignKey:AuthorID;references:ID"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ParentCommentID int       `json:"parent_comment_id" gorm:"foreignKey:ParentCommentID;references:ID;default:null"`
	PostID          int       `json:"post_id" gorm:"foreignKey:PostID;references:ID;onDelete:CASCADE"`
}
