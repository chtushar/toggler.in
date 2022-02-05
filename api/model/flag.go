package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)



type Flag struct {
	gorm.Model
	ID string `json:"id" gorm:"primary_key" validate:"required"`
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Type string `json:"type" validate:"required"`
	Value datatypes.JSON `json:"value" validate:"required"`
}
