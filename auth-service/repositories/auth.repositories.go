package repositories

import (
	entities "go-microservices/common/models"
	"net/http"

	"github.com/jinzhu/gorm"
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

func (repo *repository) Login(enitity *entities.UserEntity) (*entities.UserEntity, int) {

	db := repo.db

	var user entities.UserEntity

	checkUser := db.Select("*").Where(&entities.UserEntity{Email: enitity.Email}).First(user)

	if checkUser.RowsAffected != 0 {
		return &user, http.StatusOK
	}

	return nil, http.StatusNotFound
}

func (repo *repository) Register(enitity *entities.UserEntity) (*entities.UserEntity, int) {

	db := repo.db

	var user entities.UserEntity

	checkUser := db.Select("*").Where(&entities.UserEntity{Email: enitity.Email}).First(user)

	if checkUser.RowsAffected > 0 {
		return nil, http.StatusConflict
	}
	createUser := db.Select("*").Create(&enitity).Select(user)

	if createUser.RowsAffected > 0 {
		return nil, http.StatusExpectationFailed
	}
	
	return &user, http.StatusCreated
}
