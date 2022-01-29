package model

type Member struct {
	User User `json:"user"`
	Role string `json:"role"`
}

type Team struct {
	Base
	Name 		string `json:"name" validate:"required"`
	Members []Member	`json:"members" validate:"required"`
}