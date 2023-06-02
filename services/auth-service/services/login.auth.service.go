package services

import (
	"go-microservices/services/auth-service/models"
	entities "go-microservices/common/models"
)



func (service *service) Login(input *models.UserInputModel) (*entities.UserEntity, int) {
	userEntity := entities.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	}

	return service.repository.Login(&userEntity)
}
