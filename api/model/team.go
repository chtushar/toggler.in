package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	Base
	Name 		string    `json:"name" validate:"required"`
	Members []Member	`json:"members" validate:"required" gorm:"foreignKey:teamID"`
}

func (u *Team) BeforeCreate(tx *gorm.DB) (err error) {
  u.ID = strings.Replace(uuid.New().String(), "-", "", -1)

	return
}
