package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

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
	c.JSON(http.StatusCreated, gin.H{"data": addedLike})
}

func DeleteLike(c *gin.Context) {

	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	like := service.GetLikeById(idx)

	if like.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "like not found"})
		return
	}

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(like.UserID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	service.DeleteLike(like)

	c.JSON(http.StatusOK, gin.H{})
}

func GetLikesByPostId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	likes := service.GetLikesByPostId(idx)

	c.JSON(http.StatusOK, gin.H{"data": likes})
}
