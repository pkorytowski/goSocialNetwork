package controller

import (
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"

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

	c.JSON(http.StatusOK, gin.H{"data": addedUser})
}
