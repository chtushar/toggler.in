package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
 gorm.Model
 ID        string `sql:"primary_key"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
  b.ID = strings.Replace(uuid.New().String(), "-", "", -1)

  return
}