package main

import (
	"github.com/gin-gonic/gin"
	"socialNetwork/controller"
	"socialNetwork/model"
)

func main() {
	model.ConnectDataBase()
	r := setupRouter()
	err := r.Run()

	if err != nil {
		panic("Failed to run server!")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUserById)
	r.POST("/users", controller.AddUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	return r
}
