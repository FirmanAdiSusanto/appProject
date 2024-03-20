package data

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	PostID    string `gorm:"not null"`
	UserID    uint   `gorm:"not null"` // Menggunakan UserID sebagai kunci asing ke pengguna
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}
