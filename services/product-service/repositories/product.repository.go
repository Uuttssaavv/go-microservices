package repositories

import "github.com/jinzhu/gorm"

type Repository interface {
	//
}

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
