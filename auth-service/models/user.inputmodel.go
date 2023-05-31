package models

type UserInputModel struct {
	ID       uint   `gorm:"primary key" json:"id"`
	Email    string `gorm:"not null;unique" json:"email"`
	Phone    string `gorm:"not null;unique" json:"phone"`
	Password string `gorm:"not null" json:"-"`
}

func (input *UserInputModel) Validate() (*UserInputModel, error) {
	
	//  perform validation
	return nil, nil
}
