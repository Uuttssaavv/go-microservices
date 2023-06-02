package models

import "time"

type ProductEntity struct{
	ID uint   `gorm:"primary key"`
	Title string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price float64  `gorm:"not null"`
	Image string  `gorm:""`
	Discount float64  `gorm:"not null"`
	TotalQuanitity uint  `gorm:"not null"`
	StockQuantity uint  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

