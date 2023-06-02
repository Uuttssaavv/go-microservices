package repositories

import (
	"github.com/jinzhu/gorm"
	entities "go-microservices/common/models"
)

type Repository interface {
	Login(enitity *entities.UserEntity) (*entities.UserEntity, int)

	Register(entity *entities.UserEntity) (*entities.UserEntity, int)
}

type repository struct {
	db *gorm.DB
}

func NewAuthRepository(database *gorm.DB) *repository {

	return &repository{db: database}
}
