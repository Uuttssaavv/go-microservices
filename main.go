package main

import (
	"go-microservices/api-gateway/routes"
	"go-microservices/common/utilities"
	"log"
)

func main() {
	router := routes.SetupAppRouter()
	log.Fatal(router.Run(":" + utilities.GodotEnv("GO_PORT")))
}
