package repositories

import (
	"fmt"
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
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
	//  compare password

	checkUser := db.Select("*").Where("email=? OR phone=?", enitity.Email, enitity.Phone).First(&user)

	if checkUser.RowsAffected == 0 {
		return nil, http.StatusNotFound
	}
	
	err := utilities.ComparePassword(user.Password, enitity.Password)

	if err != nil {
		return nil, http.StatusUnauthorized
	}

	return &user, http.StatusOK

}

func (repo *repository) Register(enitity *entities.UserEntity) (*entities.UserEntity, int) {

	db := repo.db

	var user entities.UserEntity

	checkUser := db.Select("*").Where("email=? OR phone=?", enitity.Email, enitity.Phone).First(&user)
	if checkUser.RowsAffected > 0 {
		return nil, http.StatusConflict
	}
	db.NewRecord(&enitity)
	checkUser = db.Create(&enitity)

	if checkUser.Error != nil {
		fmt.Println(checkUser.Error.Error())
		return nil, http.StatusExpectationFailed
	}

	return enitity, http.StatusCreated
}
