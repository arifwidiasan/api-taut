package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminChangePass struct {
	OldPass string `json:"old_password" form:"old_password"`
	NewPass string `json:"new_password" form:"new_password"`
}
