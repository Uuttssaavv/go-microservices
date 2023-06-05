package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductEntity struct {
	ID               uint       `gorm:"primary key" json:"id"`
	Title            string     `gorm:"not null" json:"title"`
	Description      string     `gorm:"not null" json:"description"`
	Price            float64    `gorm:"not null" json:"price"`
	Image            string     `gorm:"" json:"-"`
	Discount         float64    `gorm:"not null" json:"discount"`
	TotalQuanitities uint       `gorm:"not null" json:"total_quantities"`
	StockQuantities  uint       `gorm:"not null" json:"stock_quantities"`
	UserID           uint       `gorm:"foreignkey:UserID" json:"-"`
	User             UserEntity `gorm:"foreignkey:UserID" json:"user"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (entity *ProductEntity) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	entity.UpdatedAt = time.Now().Local()

	return nil
}

func (entity *ProductEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()

	return nil
}
