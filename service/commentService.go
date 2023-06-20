package service

import (
	"socialNetwork/model"
)

func AddComment(data map[string]interface{}) (*model.Comment, error) {
	var comment model.Comment

	parentCommentId, isOk := data["parentCommentId"]

	if isOk && parentCommentId != nil {
		comment = model.Comment{
			AuthorID:        int(data["authorId"].(uint)),
			PostID:          data["postId"].(int),
			ParentCommentID: parentCommentId.(int),
			Content:         data["content"].(string),
		}
	} else {
		comment = model.Comment{
			AuthorID: int(data["authorId"].(uint)),
			PostID:   data["postId"].(int),
			Content:  data["content"].(string),
		}
	}

	if err := model.DB.Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func GetCommentsByPostId(postId int) []model.Comment {
	var comments []model.Comment
	model.DB.Where("post_id = ?", postId).Find(&comments)
	return comments
}

func GetCommentById(id int) model.Comment {
	var comment model.Comment
	model.DB.First(&comment, id)
	return comment
}

func UpdateComment(comment model.Comment) model.Comment {
	model.DB.Save(&comment)
	return comment
}

func DeleteComment(comment model.Comment) {
	model.DB.Delete(&comment)
}
