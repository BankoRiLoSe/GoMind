package model

import "time"

type KnowledgeBase struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	UUID        string    `gorm:"type:varchar(36);uniqueIndex;not null" json:"knowledge_base_id"`
	UserID      string    `gorm:"type:varchar(36);index;not null" json:"user_id"`
	Name        string    `gorm:"type:varchar(128);not null" json:"name"`
	Description string    `gorm:"type:varchar(512)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
