package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/service"
)

func RegisterUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	//check if email and password are correct
	if service.RegisterUser(email, password) {
		//if they are correct, create token
		//else return error
		c.JSON(http.StatusCreated, gin.H{})
	} else {
		c.JSON(401, gin.H{
			"message": "wrong email or password",
		})
	}
}
