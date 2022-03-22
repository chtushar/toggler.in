package models

type User struct {
	Base
	Name 			string 	 `json:"name"`
	Email 		string 	 `json:"email"`
	Password 	string
}

type AddUserParams struct {
	Name 			string 	 `json:"name" validate:"required"`
	Email 		string 	 `json:"email" validate:"required,email"`
	Password 	string 	 `json:"password" validate:"required,gte=8"`
}