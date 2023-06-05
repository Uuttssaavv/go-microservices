package repositories

import (
	entities "go-microservices/common/models"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateProduct(entities.ProductEntity) (*entities.ProductEntity, int)
	UpdateProduct(entities.ProductEntity) (*entities.ProductEntity, int)
	GetProductById(uint) (*entities.ProductEntity, int)
	GetProductsByUser(uint) (*[]entities.ProductEntity, int)
	GetAllProducts() (*[]entities.ProductEntity, int)
	DeleteProduct(uint) int
}

type repository struct {
	db *gorm.DB
}

// DeleteProduct implements Repository.
func (*repository) DeleteProduct(uint) int {
	panic("unimplemented")
}

// GetAllProducts implements Repository.
func (*repository) GetAllProducts() (*[]entities.ProductEntity, int) {
	panic("unimplemented")
}

// GetProductById implements Repository.
func (*repository) GetProductById(uint) (*entities.ProductEntity, int) {
	panic("unimplemented")
}

// GetProductsByUser implements Repository.
func (*repository) GetProductsByUser(uint) (*[]entities.ProductEntity, int) {
	panic("unimplemented")
}

// UpdateProduct implements Repository.
func (*repository) UpdateProduct(entities.ProductEntity) (*entities.ProductEntity, int) {
	panic("unimplemented")
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
