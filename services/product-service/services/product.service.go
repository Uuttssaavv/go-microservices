package services

import (
	entities "go-microservices/common/models"
	"go-microservices/services/product-service/models"
	"go-microservices/services/product-service/repositories"
)

type Service interface {
	CreateProduct(models.ProductInput) (*entities.ProductEntity, int)
	UpdateProduct(models.UpdateProductInput) (*entities.ProductEntity, int)
	GetProductById(uint) (*entities.ProductEntity, int)
	GetProductsByUser(uint) (*[]entities.ProductEntity, int)
	DeleteProduct(uint, uint) int
}
type service struct {
	repositories repositories.Repository
}

func NewProductService(repository repositories.Repository) *service {
	return &service{repositories: repository}
}
