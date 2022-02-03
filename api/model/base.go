package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
 ID        string `sql:"primary_key"`
 CreatedAt time.Time
 UpdatedAt time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
  b.ID = strings.Replace(uuid.New().String(), "-", "", -1)

  return
}