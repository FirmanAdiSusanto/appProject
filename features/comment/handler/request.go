package handler

type CommentRequest struct {
	Content string `json:"content" validate:"required"`
}
