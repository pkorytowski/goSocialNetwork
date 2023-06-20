package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

func AddFollow(c *gin.Context) {
	var input model.Follow

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idx := input.FollowsID

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(idx) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := service.FollowUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func DeleteFollow(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	follow := service.GetFollowById(idx)

	if follow.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "follow not found"})
		return
	}

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(follow.FollowsID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := service.UnfollowUser(follow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetFollowedByUserId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	follows := service.GetFollowedByUserId(idx)

	var followedUsers []model.User
	for _, follow := range follows {
		followedUsers = append(followedUsers, service.GetUserById(follow.FollowedID))
	}

	c.JSON(http.StatusOK, gin.H{"data": followedUsers})
}
