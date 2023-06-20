package service

import (
	"errors"
	"socialNetwork/model"
)

func AddLike(like model.Like) (model.Like, error) {
	var existingLike model.Like
	model.DB.Where("post_id = ? AND user_id = ?", like.PostID, like.UserID).First(&existingLike)
	if existingLike.ID != 0 {
		return existingLike, errors.New("like already exists")
	}
	model.DB.Create(&like)
	return like, nil
}

func DeleteLike(like model.Like) {
	model.DB.Delete(&like)
}

func GetLikeById(id int) model.Like {
	var like model.Like
	model.DB.First(&like, id)
	return like
}

func GetLikesByPostId(postId int) []model.Like {
	var likes []model.Like
	model.DB.Where("post_id = ?", postId).Find(&likes)
	return likes
}
