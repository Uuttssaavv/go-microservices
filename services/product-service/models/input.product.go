package models

import (
	"mime/multipart"
)

type ProductInput struct {
	ID             uint                  `gorm:"autoIncrement; primary key" json:"id" form:"id"`
	Title          string                `gorm:"not null" binding:"required" json:"title" form:"title"`
	Description    string                `gorm:"not null" binding:"required" json:"description" form:"description"`
	Price          float64               `gorm:"not null" binding:"required,numeric" json:"price" form:"price"`
	Image          *multipart.FileHeader `gorm:"" binding:"-" json:"-" form:"image"`
	Discount       float64               `gorm:"not null; default:0" binding:"omitempty,numeric" json:"discount" form:"discount"`
	TotalQuanitity uint                  `gorm:"not null" binding:"required" json:"total_quantities" form:"total_quantities"`
	StockQuantity  uint                  `gorm:"not null default:0" binding:"" json:"stock_quantities" form:"stock_quantities"`
	UserID         uint                  `gorm:"" json:"user_id"`
}
