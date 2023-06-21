package controller

import (
	"net/http"
	"socialNetwork/model"
	"socialNetwork/service"
	"socialNetwork/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users. Pass email query param to get user by email
// @Tags Users
// @Produce json
// @Param email query string false "Email"
// @Success 200 {array} model.User
// @Failure 404 "Not found"
// @Router /api/users [get]
// @Security Bearer
func GetUsers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user := service.GetUserByEmail(email)
		if user.ID == 0 {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, []model.User{user})
		return
	}
	users := service.GetUsers()
	c.JSON(http.StatusOK, users)
}

// AddUser godoc
// @Summary Add a user
// @Description Add a user. If you added user this way, you will not be able to login. Use /auth/register instead.
// @Tags Users
// @Accept json
// @Produce json
// @Param input body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 "Bad request"
// @Router /api/users [post]
// @Security Bearer
func AddUser(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedUser := service.AddUser(input)

	c.JSON(http.StatusCreated, addedUser)
}

// GetUserById godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 404 "Not found"
// @Router /api/users/{id} [get]
// @Security Bearer
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	user := service.GetUserById(idx)
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user information (name, surname, interests, hobby, age, description).
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body model.User true "User"
// @Success 200 {object} model.User
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Router /api/users/{id} [put]
// @Security Bearer
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

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @Success 204 "No content"
// @Param id path int true "User ID"
// @Router /api/users/{id} [delete]
// @Security Bearer
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	user := service.GetUserById(idx)

	service.DeleteUser(user)

	c.Status(http.StatusNoContent)
}
