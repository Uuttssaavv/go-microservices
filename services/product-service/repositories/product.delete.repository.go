package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

func (r *repository) DeleteProduct(productId uint, userId uint) int {

	db := r.db

	var product entities.ProductEntity

	err := db.Select("*").Where("id=?", productId).Find(&product).Error

	if err == nil {
		
		if product.UserID != userId {
			return http.StatusUnauthorized
		}

		err = db.Delete(&product).Error

		if err == nil {
			return http.StatusOK
		}
	}

	return http.StatusNotFound
}
