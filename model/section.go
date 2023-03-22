package model

import (
	"time"
)

type Section struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
