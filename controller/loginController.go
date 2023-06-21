package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/dto"
	"socialNetwork/service"
)

// LoginUser godoc
// @Summary Login user
// @Description Login user to get access to api
// @Tags Authorization
// @Accept json
// @Produce json
// @Param user body dto.LoginDto true "User credentials"
// @Success 200 {object} dto.TokenDto "token"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
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
	c.JSON(200, dto.TokenDto{
		Token: token,
	})

}
