package routes

import (
	"go-microservices/api-gateway/middlewares"
	"go-microservices/common/database"
	authRoute "go-microservices/services/auth-service/routes"
	productRoutes "go-microservices/services/product-service/routes"

	"github.com/gin-gonic/gin"
)
func SetupAppRouter() *gin.Engine {

	dbService := database.NewDBService()
	db := dbService.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	apiGroup := router.Group("api/v1")

	authRoute.AuthRoute(apiGroup, db)
	
	productRoute := apiGroup.Group("/product")
	
	productRoute.Use(middlewares.AuthMiddleware())

	productRoutes.ProductRoutes(productRoute, db)

	return router
}
