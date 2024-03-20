package handler

import (
	"21-api/features/user"
	"21-api/helper"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) user.UserController {
	return &controller{
		service: s,
	}
}

func (ct *controller) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}
		err = ct.service.Register(input)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "Selamat resgistrasi berhasil", nil))
	}
}

func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		var processData user.User
		processData.Hp = input.Hp
		processData.Password = input.Password

		result, token, err := ct.service.Login(processData)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		var responseData LoginResponse
		responseData.Hp = result.Hp
		responseData.Nama = result.Nama
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil login", responseData))

	}
}

func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}
		result, err := ct.service.Profile(token)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
	}
}

// Delete User
func (ct *controller) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mendapatkan token dari header
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "Token tidak valid", nil))
		}

		// Mendapatkan klaim dari token untuk mendapatkan ID pengguna
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)

		// Mendapatkan ID pengguna yang akan dihapus dari parameter URL
		targetUserID := c.Param("id")

		// Memeriksa apakah pengguna mencoba menghapus akun mereka sendiri
		if userID != targetUserID {
			return c.JSON(http.StatusForbidden,
				helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan untuk menghapus akun pengguna lain", nil))
		}

		// Memanggil service untuk menghapus pengguna
		err := ct.service.DeleteUser(targetUserID)
		if err != nil {
			// Menangani kesalahan yang terjadi saat menghapus pengguna
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "not found") {
				code = http.StatusNotFound
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		// Mengembalikan respons JSON untuk berhasil menghapus pengguna
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil menghapus pengguna", nil))
	}
}
