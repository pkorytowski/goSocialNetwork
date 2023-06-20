package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

func GetPostsByUserId(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	posts := service.GetPostsByUserId(int(userId))
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

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	post := service.GetPostById(idx)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	var input model.Post
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	post := service.GetPostById(idx)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Content = input.Content
	post = service.UpdatePost(post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	post := service.GetPostById(idx)
	service.DeletePost(post)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
