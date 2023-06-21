package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

// GetCommentsByPostId godoc
// @Summary Get comments by post id
// @Description Get comments by post id
// @Tags Comments
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Comment
// @Router /api/comments/{id} [get]
// @Security JwtAuth
func GetCommentsByPostId(c *gin.Context) {
	postId := c.Param("id")
	idx, _ := strconv.Atoi(postId)

	comments := service.GetCommentsByPostId(idx)
	c.JSON(http.StatusOK, comments)
}

// AddComment godoc
// @Summary Add comment
// @Description Add comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @StatusCreated 201 "Created"
// @StatusBadRequest 400 "Bad request"
// @Param input body model.Comment true "Comment"
// @Success 201 {object} model.Comment
// @Router /api/comments [post]
// @Security JwtAuth
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

	c.JSON(http.StatusCreated, addedComment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param input body model.Comment true "Comment"
// @Success 200 {object} model.Comment
// @Router /api/comments/{id} [put]
// @Security JwtAuth
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

	c.JSON(http.StatusOK, updatedComment)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags Comments
// @Param id path int true "Comment ID"
// @Success 204 "No content"
// @Failure 404 "Not found"
// @Failure 401 "Unauthorized"
// @Router /api/comments/{id} [delete]
// @Security JwtAuth
func DeleteComment(c *gin.Context) {
	userId, _ := token.ExtractTokenID(c)

	commentId := c.Param("id")
	idx, _ := strconv.Atoi(commentId)

	comment := service.GetCommentById(idx)

	if comment.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.AuthorID != int(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	service.DeleteComment(comment)

	c.Status(http.StatusNoContent)
}
