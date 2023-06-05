package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

func (r *repository) GetProductById(productId uint) (*entities.ProductEntity, int) {

	db := r.db

	var product entities.ProductEntity

	selectProduct := db.Select("*").Where("id=?", productId).Find(&product)
	if selectProduct.Error == nil {
		return nil, http.StatusNotFound
	}

	db.Preload("User").Find(&product)

	return &product, http.StatusFound
}
