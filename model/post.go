package model

import "time"

type Comment struct {
	CommentID       int       `json:"id" gorm:"primary_key auto_increment"`
	Author          string    `json:"author" gorm:"foreignKey:Author;references:ID"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ParentCommentID int       `json:"parent_comment_id"`
	PostID          int       `json:"post_id" gorm:"foreignKey:PostID;references:ID"`
}

//type Like struct {
//	LikeID int    `json:"id" gorm:"primary_key auto_increment"`
//	Author string `json:"author" gorm:"foreignKey:Author;references:ID"`
//	PostID int    `json:"post_id" gorm:"foreignKey:PostID;references:ID"`
//}

type Post struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	AuthorID  int       `json:"author_id" gorm:"foreignKey:AuthorID;references:ID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	//Likes    []Like    `json:"likes"`
}
