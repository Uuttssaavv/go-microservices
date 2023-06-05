package routes

import (
	"go-microservices/services/product-service/handlers"
	"go-microservices/services/product-service/repositories"
	"go-microservices/services/product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ProductRoutes(group *gin.RouterGroup, db *gorm.DB) {

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	group.POST("/create", productHandler.CreateProduct)
	group.POST("/update/:productId", productHandler.UpdateProduct)
	group.GET(":productId", productHandler.GetProductDetail)
}
