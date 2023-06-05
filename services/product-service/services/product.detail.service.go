package services

import entities "go-microservices/common/models"

func (s *service) GetProductById(productId uint) (*entities.ProductEntity, int) {

	return s.repositories.GetProductById(productId)
}
