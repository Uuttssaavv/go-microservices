package services

import (
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"go-microservices/services/product-service/models"
)

func (s *service) CreateProduct(input models.ProductInput) (*entities.ProductEntity, int) {
	var entity entities.ProductEntity

	utilities.Unmarshal(input, &entity)

	return s.repositories.CreateProduct(entity)
}
