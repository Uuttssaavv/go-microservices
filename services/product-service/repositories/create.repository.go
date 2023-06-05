package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

func (r *repository) CreateProduct(input entities.ProductEntity) (*entities.ProductEntity, int) {

	db := r.db

	db.NewRecord(&input)

	createUser := db.Create(&input)
	db.Preload("User").Find(&input)

	if createUser.Error != nil {
		return nil, http.StatusNotAcceptable
	}

	return &input, http.StatusCreated
}
