package controller

import (
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user := service.GetUserByEmail(email)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
		return
	}
	users := service.GetUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func AddUser(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedUser := service.AddUser(input)

	c.JSON(http.StatusCreated, gin.H{"data": addedUser})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	user := service.GetUserById(idx)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var input model.User
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)

	tokenId, _ := token.ExtractTokenID(c)
	if tokenId != uint(idx) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := service.GetUserById(idx)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = input.Name
	user.Surname = input.Surname
	user.Interests = input.Interests
	user.Hobby = input.Hobby
	user.Age = input.Age
	user.Description = input.Description

	updatedUser := service.UpdateUser(user)

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	user := service.GetUserById(idx)

	service.DeleteUser(user)

	c.JSON(http.StatusNoContent, gin.H{})
}
