package services

import (
	entities "go-microservices/common/models"
	"go-microservices/services/product-service/models"
	"go-microservices/services/product-service/repositories"
)

type Service interface {
	CreateProduct(models.ProductInput) (*entities.ProductEntity, int)
	UpdateProduct(models.ProductInput) (*entities.ProductEntity, int)
	GetProductById(uint) (*entities.ProductEntity, int)
	GetProductsByUser(uint) (*[]entities.ProductEntity, int)
	GetAllProducts() (*[]entities.ProductEntity, int)
	DeleteProduct(uint) int
}
type service struct {
	repositories repositories.Repository
}

func NewProductService(repository repositories.Repository) *service {
	return &service{repositories: repository}
}
func (s *service) UpdateProduct(models.ProductInput) (*entities.ProductEntity, int) { return nil, 0 }
func (s *service) GetProductById(uint) (*entities.ProductEntity, int)               { return nil, 0 }
func (s *service) GetProductsByUser(uint) (*[]entities.ProductEntity, int)          { return nil, 0 }
func (s *service) GetAllProducts() (*[]entities.ProductEntity, int)                 { return nil, 0 }
func (s *service) DeleteProduct(uint) int                                           { return 0 }
