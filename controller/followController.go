package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"
)

// AddFollow godoc
// @Summary Add follow
// @Description Follow an user
// @Tags Follow
// @Accept json
// @Param input body model.Follow true "Follow"
// Success 201 "Created"
// @Router /api/follows [post]
// @Security Bearer
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
	c.Status(http.StatusNoContent)
}

// DeleteFollow godoc
// @Summary Delete follow
// @Description Delete follow
// @Tags Follow
// @Param id path int true "Follow ID"
// @Success 204 "No Content"
// @Router /api/follows/{id} [delete]
// @Security Bearer
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

	c.Status(http.StatusNoContent)
}

// GetFollowedByUserId godoc
// @Summary Get followed by user id
// @Description Get followed by user id
// @Tags Follow
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {array} model.User
// @Router /api/follows/{id} [get]
// @Security Bearer
func GetFollowedByUserId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	follows := service.GetFollowedByUserId(idx)

	followedUsers := make([]model.User, 0)
	for _, follow := range follows {
		followedUsers = append(followedUsers, service.GetUserById(follow.FollowedID))
	}

	c.JSON(http.StatusOK, followedUsers)
}

// GetFollowsByUserId godoc
// @Summary Get follows by user id
// @Description Get follows by user id
// @Tags Users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {array} model.User
// @Router /api/users/{id}/followers [get]
// @Security Bearer
func GetFollowsByUserId(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	follows := service.GetFollowsByUserId(idx)

	followedUsers := make([]model.User, 0)
	for _, follow := range follows {
		followedUsers = append(followedUsers, service.GetUserById(follow.FollowsID))
	}

	c.JSON(http.StatusOK, followedUsers)
}
