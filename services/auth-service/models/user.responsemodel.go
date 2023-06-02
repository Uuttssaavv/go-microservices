package models


type UserResponse struct {
	ID        uint   `gorm:"primary key" json:"id"`
	Name      string `gorm:"" json:"name"`
	Email     string `gorm:"not null;unique" json:"email"`
	Image     string `gorm:"not null;unique" json:"image"`
	Phone     string `gorm:"not null;unique" json:"phone"`
	Password  string `gorm:"not null" json:"-"`
	Address   string `gorm:"" json:"address"`
	Token     string `json:"token"`
}
