package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
