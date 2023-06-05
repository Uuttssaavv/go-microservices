package models

import (
	"go-microservices/common/utilities"
	"time"

	"github.com/jinzhu/gorm"
)

type UserEntity struct {
	ID        uint   `gorm:"primary key" json:"id"`
	Name      string `gorm:"" json:"name"`
	Email     string `gorm:"not null;unique" json:"email"`
	Image     string `gorm:"not null" json:"image"`
	Phone     string `gorm:"not null;unique" json:"phone"`
	Password  string `gorm:"not null" json:"-"`
	Address   string `gorm:"" json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *UserEntity) BeforeCreate(*gorm.DB) error {
	//  encrypt the password
	entity.Password = utilities.HashPassword(entity.Password)

	//  update the createdAt field
	entity.CreatedAt = time.Now().Local()
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (entity *UserEntity) BeforeUpdate(*gorm.DB) error {

	//  update the UpdatedAt field
	entity.UpdatedAt = time.Now().Local()

	return nil
}
