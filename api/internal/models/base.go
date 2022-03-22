package models

import "time"

type Base struct {
	ID 				uint32 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}