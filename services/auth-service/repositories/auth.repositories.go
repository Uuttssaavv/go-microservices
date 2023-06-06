package repositories

import (
	"github.com/jinzhu/gorm"
	entities "go-microservices/common/models"
)
//go:generate mockery --name=Repository --output=../tests/mock/ --case=underscore --with-expecter
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
