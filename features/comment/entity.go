package comment

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentController interface {
	AddComment() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}

type CommentModel interface {
	DeleteComment(userID string, postID uint, AddComment string) error
	AddComment(userID string, AddComment string) error
}

type CommentService interface {
	AddComment(token *jwt.Token, postID uint, AddComment string) error
	DeleteComment(token *jwt.Token, postID uint, commentID string) error
}

type Comment struct {
	Id        uint
	PostID    uint
	UserID    string
	Content   string
	CreatedAt time.Time
}

type AddCommentComment struct {
	postID  uint
	Content string
}
