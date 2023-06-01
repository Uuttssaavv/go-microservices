package models

type UserInputModel struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"len=10"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

func (input *UserInputModel) Validate() (*UserInputModel, error) {
	
	//  perform validation
	return nil, nil
}
