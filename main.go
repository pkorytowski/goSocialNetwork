package main

import (
	"github.com/gin-gonic/gin"
	"socialNetwork/controller"
	"socialNetwork/model"
)

func main() {
	r := gin.Default()

	model.ConnectDataBase()

	r.GET("/users", controller.GetUsers)
	r.POST("/users", controller.AddUser)
	r.Run()
}
