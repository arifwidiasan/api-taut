package model

import (
	"time"
)

type Sosmed struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Facebook  string    `json:"facebook"`
	Twitter   string    `json:"twitter"`
	Instagram string    `json:"instagram"`
	Linkedin  string    `json:"linkedin"`
	Github    string    `json:"github"`
	Telegram  string    `json:"telegram"`
	Youtube   string    `json:"youtube"`
	Tiktok    string    `json:"tiktok"`
	Line      string    `json:"line"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
