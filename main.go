package main

import (
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
	dbService.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	apiGroup := router.Group("api/v1")
	
	logrus.Info("ApiGroup=>", apiGroup)

	return router
}
