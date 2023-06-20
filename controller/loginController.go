package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/dto"
	"socialNetwork/service"
)

func LoginUser(c *gin.Context) {
	loginDto := dto.LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check if email and password are correct
	token, err := service.LoginUser(loginDto)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "wrong email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})

}
