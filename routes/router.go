package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Adicione sua origem
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"}, // Métodos permitidos
        AllowHeaders:     []string{"Content-Type", "Authorization"}, // Permite cabeçalho Authorization
        ExposeHeaders:    []string{"Authorization"},
        AllowCredentials: true,
    }))

	InitializeRoutes(r)

	r.Run(":8080")
}