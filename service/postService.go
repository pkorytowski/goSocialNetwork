package service

import (
	"errors"
	"socialNetwork/model"
)

func GetPostsByUserId(userId int) []model.Post {
	var posts []model.Post
	model.DB.Where("author_id = ?", userId).Find(&posts)
	return posts
}

func GetPostById(id int) (*model.Post, error) {
	var post model.Post
	model.DB.First(&post, id)

	if post.ID == 0 {
		return nil, errors.New("post not found")
	}

	return &post, nil
}

func AddPost(data map[string]interface{}) (*model.Post, error) {
	post := model.Post{
		AuthorID: int(data["userId"].(uint)),
		Content:  data["content"].(string),
	}

	if err := model.DB.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePost(post *model.Post) *model.Post {
	model.DB.Save(post)
	return post
}

func DeletePost(post *model.Post) {
	model.DB.Delete(post)
}
