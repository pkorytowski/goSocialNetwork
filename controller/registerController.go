package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/dto"
	"socialNetwork/service"
)

func RegisterUser(c *gin.Context) {
	loginDto := dto.LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check if email and password are correct

	user := service.GetUserByEmail(loginDto.Email)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists, use another email.",
		})
		return
	}

	if service.RegisterUser(loginDto) {
		//if they are correct, create token
		//else return error
		c.JSON(http.StatusCreated, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong email or password.",
		})
	}
}
