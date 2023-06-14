package controller

import (
	"net/http"
	"socialNetwork/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	model.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func AddUser(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	model.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}
