package models

type ProductInput struct {
	ID             uint    `gorm:"primary key" json:"id"`
	Title          string  `gorm:"not null" binding:"required" json:"title"`
	Description    string  `gorm:"not null" binding:"required" json:"description"`
	Price          float64 `gorm:"not null" binding:"required,numeric" json:"price"`
	Image          string  `gorm:"" binding:"required" json:"image"`
	Discount       float64 `gorm:"not null; default:0" binding:"omitempty,numeric" json:"discount"`
	TotalQuanitity uint    `gorm:"not null" binding:"" json:"total_quantities"`
	StockQuantity  uint    `gorm:"not null" binding:"" json:"stock_quantities"`
}
