package routes

import (
	"go-microservices/auth-service/handlers"
	"go-microservices/auth-service/repositories"
	"go-microservices/auth-service/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AuthRoute(group *gin.RouterGroup, db *gorm.DB) {

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	group.POST("/login", authHandler.LoginHandler)
	group.POST("/register", authHandler.RegisterHandler)
}
