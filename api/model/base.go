package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
 gorm.Model
 ID        string `sql:"primary_key"`
 CreatedAt time.Time `json:"createdAt"`
 UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.ID = strings.Replace(uuid.New().String(), "-", "", -1)

	validity, err := u.IsValid()
  if !validity {
    fmt.Print("Not a valid user model", err.Error())
  }
  return
}