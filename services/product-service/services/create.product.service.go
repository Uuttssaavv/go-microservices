package services

import (
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"go-microservices/services/product-service/models"
)

func (s *service) CreateProduct(input models.ProductInput) (*entities.ProductEntity, int) {
	var entity entities.ProductEntity

	utilities.Unmarshal(input, &entity)
	//  assign the stock quantities as total quantities when creating new product
	entity.StockQuantities = input.TotalQuanitity
	entity.UserID = input.UserID
	return s.repositories.CreateProduct(entity)
}
