package model

import "time"

type User struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
