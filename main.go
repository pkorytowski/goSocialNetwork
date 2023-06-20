package main

import (
	"github.com/gin-gonic/gin"
	"socialNetwork/controller"
	"socialNetwork/middleware"
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

	public := r.Group("/auth")

	//register
	public.POST("/register", controller.RegisterUser)

	//login
	public.POST("/login", controller.LoginUser)

	protected := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())

	//users
	protected.GET("/users", controller.GetUsers)
	protected.GET("/users/:id", controller.GetUserById)
	protected.POST("/users", controller.AddUser)
	protected.PUT("/users/:id", controller.UpdateUser)
	protected.DELETE("/users/:id", controller.DeleteUser)

	//posts
	protected.GET("/posts", controller.GetPostsByUserId)
	protected.POST("/posts", controller.AddPost)
	protected.GET("/posts/:id", controller.GetPostById)
	protected.PATCH("/posts/:id", controller.UpdatePost)
	protected.DELETE("/posts/:id", controller.DeletePost)

	// comments
	protected.GET("/comments/:id", controller.GetCommentsByPostId)
	protected.POST("/comments", controller.AddComment)
	protected.PATCH("/comments/:id", controller.UpdateComment)
	protected.DELETE("/comments/:id", controller.DeleteComment)
	return r
}
