package routes

import (
	"go-microservices/common/middlewares"
	"go-microservices/services/product-service/handlers"
	"go-microservices/services/product-service/repositories"
	"go-microservices/services/product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ProductRoutes(group *gin.RouterGroup, db *gorm.DB) {
	// using auth middleware to verify the
	// user token
	group.Use(middlewares.AuthMiddleware())

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	group.POST("/create", productHandler.CreateProduct)
}
