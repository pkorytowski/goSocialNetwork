package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

// AddLike godoc
// @Summary Add a like
// @Description Add a like to a post
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param input body model.Like true "Like"
// @Success 201 {object} model.Like
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Router /api/likes [post]
// @Security Bearer
func AddLike(c *gin.Context) {
	var input model.Like

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idx := input.UserID

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(idx) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	addedLike, err := service.AddLike(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedLike)
}

// DeleteLike godoc
// @Summary Delete a like
// @Description Delete a like
// @Tags Likes
// @Success 204 "No Content"
// @Failure 404 "Not Found"
// @Failure 401 "Unauthorized"
// @Param id path int true "Like ID"
// @Router /api/likes/{id} [delete]
// @Security Bearer
func DeleteLike(c *gin.Context) {

	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	like := service.GetLikeById(idx)

	if like.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(like.UserID) {
		c.Status(http.StatusUnauthorized)
		return
	}

	service.DeleteLike(like)

	c.Status(http.StatusNoContent)
}

// GetLikesByPostId godoc
// @Summary Get likes by post id
// @Description Get likes for a post
// @Tags Likes
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {array} model.Like
// @Router /api/posts/{id}/likes [get]
// @Security Bearer
func GetLikesByPostId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	likes := service.GetLikesByPostId(idx)

	c.JSON(http.StatusOK, likes)
}
