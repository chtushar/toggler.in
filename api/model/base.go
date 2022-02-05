package model

import (
	"time"
)

type Base struct {
	ID        string 			`json:"id" gorm:"primary_key;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
