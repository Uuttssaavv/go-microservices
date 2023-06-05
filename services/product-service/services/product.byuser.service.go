package services

import entities "go-microservices/common/models"

func (s *service) GetProductsByUser(userId uint) (*[]entities.ProductEntity, int) {

	return s.repositories.GetProductsByUser(userId)
}
