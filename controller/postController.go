package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

func GetPostsByUserId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	posts := service.GetPostsByUserId(idx)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

type PostInput struct {
	Content string `json:"content"`
}

func AddPost(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postInput := map[string]interface{}{
		"userId":  userId,
		"content": input.Content,
	}

	addedPost, err := service.AddPost(postInput)

	c.JSON(http.StatusCreated, gin.H{"data": addedPost})
}
