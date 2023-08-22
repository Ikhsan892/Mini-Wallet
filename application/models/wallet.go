package models

import "time"

type Wallet struct {
	BaseModel
	OwnedBy   string     `json:"owned_by" gorm:"not null"`
	Status    string     `json:"status"`
	EnabledAt *time.Time `json:"enabled_at"`
	Balance   float64    `json:"balance"`
}
