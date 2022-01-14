package model

import "toggler.in/api/utils"

type User struct {
	Base
	FirstName string 			`json:"firstName" validate:"required"`
	LastName 	string 			`json:"lastName" validate:"required"`
	Email    	string 			`json:"email" validate:"required,email"`
	Password 	string 			`json:"password" validate:"min=8"`
}

func (u *User) IsValid() (bool, error) {
	err := utils.Validate.Struct(u)
	return err == nil, err
}