package main

import (
	"go-microservices/auth-service/routes"
	"go-microservices/common/database"
	"go-microservices/common/utilities"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	routes.AuthRoute(apiGroup, db)

	logrus.Info("ApiGroup=>", apiGroup)

	return router
}
