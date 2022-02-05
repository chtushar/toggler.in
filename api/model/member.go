package model

import "gorm.io/gorm"

type Role string

const (
	admin Role = "admin"
	member Role = "member"
)

type Member struct {
	gorm.Model
	User User `json:"user"`
	Role Role `json:"role"`
	TeamID string `json:"teamID"`
}