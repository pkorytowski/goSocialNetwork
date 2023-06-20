package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

func GetCommentsByPostId(c *gin.Context) {
	postId := c.Param("id")
	idx, _ := strconv.Atoi(postId)

	comments := service.GetCommentsByPostId(idx)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func AddComment(c *gin.Context) {
	userId, _ := token.ExtractTokenID(c)

	var input model.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentInput := map[string]interface{}{
		"postId":          input.PostID,
		"content":         input.Content,
		"authorId":        userId,
		"parentCommentId": input.ParentCommentID,
	}

	addedComment, err := service.AddComment(commentInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": addedComment})
}

func UpdateComment(c *gin.Context) {
	userId, _ := token.ExtractTokenID(c)

	commentId := c.Param("id")
	idx, _ := strconv.Atoi(commentId)

	comment := service.GetCommentById(idx)

	if comment.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
		return
	}

	if comment.AuthorID != int(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input model.Comment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = input.Content

	updatedComment := service.UpdateComment(comment)

	c.JSON(http.StatusOK, gin.H{"data": updatedComment})
}

func DeleteComment(c *gin.Context) {
	userId, _ := token.ExtractTokenID(c)

	commentId := c.Param("id")
	idx, _ := strconv.Atoi(commentId)

	comment := service.GetCommentById(idx)

	if comment.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
		return
	}

	if comment.AuthorID != int(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	service.DeleteComment(comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
