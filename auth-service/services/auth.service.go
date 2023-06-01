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