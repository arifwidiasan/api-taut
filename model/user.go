package model

import "time"

type User struct {
	ID                     int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username               string    `gorm:"uniqueIndex;not null;size:32" json:"username"`
	Password               string    `gorm:"not null" json:"password"`
	Email                  string    `gorm:"unique;not null" json:"email"`
	Name                   string    `gorm:"not null" json:"name"`
	Job                    string    `gorm:"not null" json:"job"`
	BornDate               time.Time `gorm:"not null" json:"born_date"`
	PhoneNumber            string    `gorm:"not null" json:"phone_number"`
	Address                string    `json:"address"`
	About                  string    `json:"about"`
	ProfilePicturePathFile string    `json:"profile_picture_path_file"`
	QrcodePathFile         string    `json:"qrcode_path_file"`
	CreatedAt              time.Time `json:"created_at"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserChangePass struct {
	OldPass string `json:"old_password" form:"old_password"`
	NewPass string `json:"new_password" form:"new_password"`
}
