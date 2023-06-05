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
	DeleteProduct(uint,uint) int
}

type repository struct {
	db *gorm.DB
}

// GetAllProducts implements Repository.
func (*repository) GetAllProducts() (*[]entities.ProductEntity, int) {
	panic("unimplemented")
}

// GetProductsByUser implements Repository.
func (*repository) GetProductsByUser(uint) (*[]entities.ProductEntity, int) {
	panic("unimplemented")
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
