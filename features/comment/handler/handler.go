package handler

import (
	comment "21-api/features/comment"
	"21-api/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s comment.CommentService
}

func NewHandler(service comment.CommentService) comment.CommentController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CommentRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var inputProcess comment.Comment
		inputProcess.Content = input.Content
		result, err := ct.s.AddComment(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Komentar ditambahkan", result))
	}
}

// Hapus Komentar
func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Membaca commentID dari parameter URL
		commentID := c.Param("commentID")

		// Konversi commentID ke tipe data uint
		id, err := strconv.ParseUint(commentID, 10, 64)
		if err != nil {
			log.Println("error parsing commentID:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		// Memanggil service untuk menghapus komentar
		err = ct.s.DeleteComment(uint(id))
		if err != nil {
			log.Println("error delete db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		// Mengembalikan respons sukses
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Komentar dihapus", nil))
	}
}

// Tambah Komentar
func (ct *controller) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CommentRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var inputProcess comment.Comment
		inputProcess.Content = input.Content
		result, err := ct.s.AddComment(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Komentar ditambahkan", result))
	}
}
