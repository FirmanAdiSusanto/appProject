package services

import (
	"21-api/features/comment"
	"21-api/helper"
	"21-api/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m comment.CommentModel
	v *validator.Validate
}

func NewCommentService(model comment.CommentModel) comment.CommentService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

// Fungsi untuk menambahkan komentar baru ke dalam database
func (s *service) AddComment(token *jwt.Token, postID uint, content string) error {
	decodeUserID := middlewares.DecodeToken(token)
	if decodeUserID == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.AddComment(decodeUserID, content)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

// Fungsi untuk menghapus komentar dari database
func (s *service) DeleteComment(token *jwt.Token, postID uint, commentID string) error {
	decodeUserID := middlewares.DecodeToken(token)
	if decodeUserID == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.DeleteComment(decodeUserID, postID, commentID) // Hapus postID dari sini karena tidak diperlukan
	if err != nil {
		return err
	}

	return nil
}
