package services

import (
	"21-api/features/post"
	"21-api/helper"
	"21-api/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m post.PostModel
	v *validator.Validate
}

func NewPostService(model post.PostModel) post.PostService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddPost(userid *jwt.Token, pictureBaru string, contentBaru string) (post.Post, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return post.Post{}, errors.New("data tidak valid")
	}

	// Validasi pictureBaru dan contentBaru hanya jika keduanya tidak kosong
	if pictureBaru != "" && contentBaru != "" {
		err := s.v.Var(pictureBaru, "required")
		if err != nil {
			log.Println("error validasi judul", err.Error())
			return post.Post{}, err
		}

		err = s.v.Var(contentBaru, "required")
		if err != nil {
			log.Println("error validasi deskripsi", err.Error())
			return post.Post{}, err
		}
	}

	result, err := s.m.AddPost(id, pictureBaru, contentBaru)
	if err != nil {
		return post.Post{}, errors.New(helper.ErrorGeneralServer)
	}

	return result, nil
}

func (s *service) UpdatePost(userid *jwt.Token, postID uint, data post.Post) (post.Post, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return post.Post{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(data)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return post.Post{}, err
	}

	result, err := s.m.UpdatePost(id, postID, data)
	if err != nil {
		return post.Post{}, errors.New(helper.ErrorGeneralServer)
	}

	return result, nil
}

func (s *service) DeletePost(userid *jwt.Token, postID uint) error {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.DeletePost(postID)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) GetAllPosts() ([]post.Post, error) {
	posts, err := s.m.GetAllPosts() // 0 digunakan untuk menunjukkan bahwa kita tidak memerlukan userID
	if err != nil {
		return nil, errors.New(helper.ErrorGeneralServer)
	}

	return posts, nil
}

func (s *service) GetPostByID(postID uint) (*post.Post, error) {
	post, err := s.m.GetPostByID(postID)
	if err != nil {
		return nil, errors.New(helper.ErrorGeneralServer)
	}
	return post, nil
}
