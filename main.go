package main

import (
	"go-microservices/common/database"
	"go-microservices/common/utilities"
	authRoute "go-microservices/services/auth-service/routes"
	productRoute "go-microservices/services/product-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupAppRouter()
	log.Fatal(router.Run(":" + utilities.GodotEnv("GO_PORT")))

}

func SetupAppRouter() *gin.Engine {

	dbService := database.NewDBService()
	db := dbService.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	apiGroup := router.Group("api/v1")

	authRoute.AuthRoute(apiGroup, db)
	productRoute.ProductRoutes(apiGroup, db)

	return router
}
