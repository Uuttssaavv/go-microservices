package repositories

import (
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"net/http"
)

func (repo *repository) Login(enitity *entities.UserEntity) (*entities.UserEntity, int) {

	db := repo.db

	var user entities.UserEntity
	
	checkUser := db.Select("*").Where("email=? OR phone=?", enitity.Email, enitity.Phone).First(&user)

	if checkUser.RowsAffected == 0 {
		return nil, http.StatusNotFound
	}
	//  compare password
	err := utilities.ComparePassword(user.Password, enitity.Password)

	if err != nil {
		return nil, http.StatusUnauthorized
	}

	return &user, http.StatusOK

}
