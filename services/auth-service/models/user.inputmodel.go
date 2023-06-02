package models

type UserInputModel struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty,numeric,len=10"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}
