package comment

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentController interface {
	AddComment() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentModel interface {
	InsertComment(userID string, contentBaru Comment) (Comment, error)
	DeleteComment(commentID uint) error
	//GetComment(userID string) ([]Comment, error)
	AddComment(userID *jwt.Token, contentBaru Comment) (Comment, error)
}

type CommentService interface {
	AddComment(userID *jwt.Token, contentBaru Comment) (Comment, error)
	DeleteComment(commentID uint) error
}

type Comment struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
