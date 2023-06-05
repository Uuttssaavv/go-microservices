package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

func (r *repository) GetProductsByUser(userId uint) (*[]entities.ProductEntity, int) {

	db := r.db

	var products []entities.ProductEntity
	
	
	selectProduct := db.Preload("User").Where("user_id = ?", userId).Find(&products)
	
	if selectProduct.Error != nil {
		return nil, http.StatusNotFound
	}
	
	return &products, http.StatusFound
}
