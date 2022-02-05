package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"toggler.in/api/utils/validator"
)

type User struct {
	Base
	FirstName string 			`json:"firstName" validate:"required"`
	LastName 	string 			`json:"lastName" validate:"required"`
	Email    	string 			`json:"email" gorm:"unique" validate:"required,email"`
	Password 	string 			`json:"-" validate:"min=8"`
}

func (u *User) IsValid() (bool, error) {
	err := validator.Validate.Struct(u)
	return err == nil, err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.ID = strings.Replace(uuid.New().String(), "-", "", -1)

	validity, err := u.IsValid()
	if !validity {
		fmt.Println(err.Error())
    return errors.New("User is not valid")
  }
	return nil
}
