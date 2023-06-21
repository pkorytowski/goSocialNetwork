package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/dto"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

// GetPostsByUserId godoc
// @Summary Get posts by user id
// @Description Get posts for given user
// @Tags Posts
// @Produces json
// @Param id path int true "User ID"
// @Success 200 {array} model.Post
// @Router /api/users/{id}/posts [get]
// @Security JwtAuth
func GetPostsByUserId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	posts := service.GetPostsByUserId(idx)
	c.JSON(http.StatusOK, posts)
}

// AddPost godoc
// @Summary Add a post
// @Description Add a post
// @Tags Posts
// @Accept json
// @Produce json
// @Param input body dto.PostInputDto true "Post"
// @Success 201 {object} model.Post
// @Failure 400 "Bad request"
// @Router /api/posts [post]
// @Security JwtAuth
func AddPost(c *gin.Context) {
	userId, _ := token.ExtractTokenID(c)

	var input dto.PostInputDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postInput := map[string]interface{}{
		"userId":  userId,
		"content": input.Content,
	}

	addedPost, _ := service.AddPost(postInput)

	c.JSON(http.StatusCreated, addedPost)
}

// GetPostById godoc
// @Summary Get post by id
// @Description Get post by id
// @Tags Posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post
// @Router /api/posts/{id} [get]
// @Security JwtAuth
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	post := service.GetPostById(idx)
	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update a post
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param input body model.Post true "Post"
// @Success 200 {object} model.Post
// @Failure 400 "Bad request"
// @Router /api/posts/{id} [put]
// @Security JwtAuth
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
	c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post
// @Tags Posts
// @Success 204 "No Content"
// @Failure 404 "Not Found"
// @Param id path int true "Post ID"
// @Router /api/posts/{id} [delete]
// @Security JwtAuth
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	post := service.GetPostById(idx)
	service.DeletePost(post)
	c.Status(http.StatusNoContent)
}
