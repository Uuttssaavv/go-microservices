package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

func (r *repository) UpdateProduct(input entities.ProductEntity) (*entities.ProductEntity, int) {

	db := r.db
	var product entities.ProductEntity
	updateProduct := db.Select("*").Where("id=? AND user_id=?", input.ID, input.UserID).Find(&product)

	if updateProduct.Error != nil {
		return nil, http.StatusNotFound
	}
	updateProduct = db.Model(&input).Updates(&input)

	if updateProduct.Error != nil {
		return nil, http.StatusForbidden
	}
	
	db.Preload("User").Find(&input)

	return &input, http.StatusOK
}
