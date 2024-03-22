package data

import (
	comment "21-api/features/comment/data"
	post "21-api/features/post/data"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment"`
	name      string `validate:"required"`
	Email     string `gorm:"unique"`
	Password  string
	Handphone string
	Posts     []post.Post       `gorm:"foreignKey:UserID"`
	Comments  []comment.Comment `gorm:"foreignKey:UserID"`
}
