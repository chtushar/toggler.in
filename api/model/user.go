package model

import (
	"fmt"

	"gorm.io/gorm"
	"toggler.in/api/utils/validator"
)

type User struct {
	Base
	FirstName string 			`json:"firstName" validate:"required"`
	LastName 	string 			`json:"lastName" validate:"required"`
	Email    	string 			`json:"email" validate:"required,email"`
	Password 	string 			`json:"-" validate:"min=8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validity, err := u.IsValid()
	if !validity {
    fmt.Print("Not a valid user model", err.Error())
  }
	return
}

func (u *User) IsValid() (bool, error) {
	err := validator.Validate.Struct(u)
	return err == nil, err
}