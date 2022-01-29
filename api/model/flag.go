package model

import "gorm.io/datatypes"

type Flag struct {
	Base
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Type string `json:"type" validate:"required"`
	Value datatypes.JSON `json:"value" validate:"required"`
}