package controller

import (
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
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
	user := service.GetUserById(idx)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = input.Name
	user.Surname = input.Surname
	user.Email = input.Email
	user.Password = input.Password
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
