package repositories

import (
	entities "go-microservices/common/models"
	"net/http"
)

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
		return nil, http.StatusExpectationFailed
	}

	return enitity, http.StatusCreated
}
