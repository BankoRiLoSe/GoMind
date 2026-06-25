package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	UUID      string    `gorm:"type:varchar(36);uniqueIndex;not null" json:"user_id"`
	Username  string    `gorm:"type:varchar(64);uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
