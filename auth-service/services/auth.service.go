package services

import (
	"go-microservices/auth-service/models"
	"go-microservices/auth-service/repositories"
	entities "go-microservices/common/models"
)

type Service interface {
	Login(*models.UserInputModel) (*entities.UserEntity, int)

	Register(*models.UserInputModel) (*entities.UserEntity, int)
}

type service struct {
	repository repositories.Repository
}

func NewAuthService(repo repositories.Repository) *service {
	return &service{repository: repo}
}

func (service *service) Login(input *models.UserInputModel) (*entities.UserEntity, int) {
	userEntity := entities.UserEntity{
		Email:    input.Email,
		Password: input.Password,
	}

	return service.repository.Login(&userEntity)
}

func (service *service) Register(input *models.UserInputModel) (*entities.UserEntity, int) {
	userEntity := entities.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	}

	return service.repository.Register(&userEntity)
}
