package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"socialNetwork/controller"
	_ "socialNetwork/docs"
	"socialNetwork/middleware"
	"socialNetwork/model"
)

//@title Social Network API
//@version 1.0
//@description This is a social network API server created for GO classes.
//@host localhost:8080
//@BasePath /api

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

	docs := r.Group("/docs")
	docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())

	//users
	protected.GET("/users", controller.GetUsers)
	protected.GET("/users/:id", controller.GetUserById)
	protected.GET("/users/:id/posts", controller.GetPostsByUserId)
	protected.POST("/users", controller.AddUser)
	protected.PUT("/users/:id", controller.UpdateUser)
	protected.DELETE("/users/:id", controller.DeleteUser)

	//posts
	protected.POST("/posts", controller.AddPost)
	protected.GET("/posts/:id", controller.GetPostById)
	protected.PATCH("/posts/:id", controller.UpdatePost)
	protected.DELETE("/posts/:id", controller.DeletePost)
	protected.GET("/posts/:id/likes", controller.GetLikesByPostId)
	protected.GET("/posts/:id/comments", controller.GetCommentsByPostId)

	//likes
	protected.POST("/likes", controller.AddLike)
	protected.DELETE("/likes/:id", controller.DeleteLike)

	// comments
	protected.POST("/comments", controller.AddComment)
	protected.PATCH("/comments/:id", controller.UpdateComment)
	protected.DELETE("/comments/:id", controller.DeleteComment)
	return r
}
