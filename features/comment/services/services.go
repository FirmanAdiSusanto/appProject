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

// Fungssi digunakan untuk menambahkan komentar baru ke dalam database
// func (s *service) AddComment(pemilik *jwt.Token, commentBaru comment.Comment) (comment.Comment, error) {
// 	hp := middlewares.DecodeToken(pemilik) //Melakukan dekode token akses JWT
// 	if hp == "" {
// 		log.Println("error decode token:", "token tidak ditemukan")
// 		return comment.Comment{}, errors.New("data tidak valid")
// 	} //Jika ada kesalahan validasi, fungsi akan mengembalikan error bersamaan dengan pesan kesalahan

// 	err := s.v.Struct(&commentBaru)
// 	if err != nil {
// 		log.Println("error validasi", err.Error())
// 		return comment.Comment{}, err
// 	}

// 	result, err := s.m.InsertComment(hp, commentBaru) //Menyimpan data komentar baru ke dalam database
// 	if err != nil {
// 		return comment.Comment{}, errors.New(helper.ServerGeneralError)
// 	}

// 	return result, nil
// }

// Fungsi untuk menghapus komentar dari database
func (s *service) DeleteComment(commentID uint) error {
	// Memanggil method DeleteComment dari model untuk menghapus komentar
	err := s.m.DeleteComment(commentID)
	if err != nil {
		log.Println("error delete comment:", err.Error())
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}

// Fungsi untuk menambahkan komentar baru ke dalam database
func (s *service) AddComment(userID *jwt.Token, contentBaru comment.Comment) (comment.Comment, error) {
	// Mendekode token akses JWT untuk mendapatkan userID
	hp := middlewares.DecodeToken(userID)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return comment.Comment{}, errors.New("data tidak valid")
	}

	// Validasi struktur komentar baru
	err := s.v.Struct(&contentBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return comment.Comment{}, err
	}

	// Memasukkan komentar baru ke dalam database menggunakan model
	result, err := s.m.InsertComment(hp, contentBaru)
	if err != nil {
		return comment.Comment{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}
