package model

type Role string
const (
	admin Role = "admin"
	member Role = "member"
)
type Member struct {
	User User `json:"user"`
	Role Role `json:"role"`
}

type Team struct {
	Base
	Name 		string `json:"name" validate:"required"`
	Members []Member	`json:"members" validate:"required"`
	Flag 		[]Flag	`json:"flag"`
}