package controllers

import (
	"net/http"

	"github.com/djuniordev/SAST-MA/database"
	"github.com/djuniordev/SAST-MA/repositories"
	"github.com/djuniordev/SAST-MA/services"
	"github.com/gin-gonic/gin"
)

func ListUsers(ctx *gin.Context) {

	userRepository := &repositories.UserRepository{Db: database.DB}
	UserService := &services.UserService{UserRepository: userRepository}

	usersDto, err := UserService.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, usersDto)
}
