package data

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null"`
	UserID    string `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}
