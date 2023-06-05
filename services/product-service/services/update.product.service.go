package services

import (
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"go-microservices/services/product-service/models"
)

func (s *service) UpdateProduct(input models.UpdateProductInput) (*entities.ProductEntity, int) {
	var product entities.ProductEntity

	utilities.Unmarshal(input, &product)
	product.UserID= input.UserID
	
	return s.repositories.UpdateProduct(product)
}
