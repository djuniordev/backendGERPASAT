package controllers

import (
	"net/http"

	"github.com/djuniordev/SAST-MA/database"
	"github.com/djuniordev/SAST-MA/dto"
	"github.com/djuniordev/SAST-MA/repositories"
	"github.com/djuniordev/SAST-MA/services"
	"github.com/djuniordev/SAST-MA/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var login dto.LoginDto

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados de entrada inválidos",
		})
		return
	}

	if err := dto.ValidaDadosDeLogin(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	loginRepository := &repositories.LoginRepository{Db: database.DB}
	loginService := &services.LoginService{LoginRepository: loginRepository}

	userDto, err := loginService.Execute(&login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Gera o token JWT
	accessToken, err := utils.GenerateToken(*userDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(*userDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate refresh token"})
		return
	}

	// Retorna o token
	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})

}
