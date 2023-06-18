package controller

import (
	"github.com/gin-gonic/gin"
	"socialNetwork/service"
)

func LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	//check if email and password are correct
	token, err := service.LoginUser(email, password)

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
