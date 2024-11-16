package routes

import (
	"github.com/djuniordev/SAST-MA/controllers"
	"github.com/djuniordev/SAST-MA/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.GET("/ola", middleware.AuthMiddleware(), controllers.HelloWorld)
	r.GET("/user", middleware.AuthMiddleware(), controllers.ListUsers)
}
