package data

import (
	"21-api/features/comment/data"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Picture string
	Content string
	Comment []data.Comment `gorm:"foreignKey:PostID"`
}
