package routes

import (
	"TeraApps/config"
	user "TeraApps/features/user"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController) { //, cc comment.CommentController) {
	userRoute(c, ctl)

}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/login", ctl.Login())

	//Register User
	c.POST("/register", ctl.RegisterUser()) //Endpoint untuk API

	//GetUserByHp

	c.GET("/users/:hp", ctl.GetUserByHP(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

	//DeleteUser
	c.DELETE("/users/:id", ctl.DeleteUser(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

	// UpdateUser
	c.PUT("/users/:hp", ctl.UpdateUser(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

}

// func commentRoute(c *echo.Echo, cc comment.CommentController) {
// 	//Menambahkan Komentar
// 	c.POST("/comment", cc.AddComment())

// 	//Delete Komentar
// 	c.DELETE("/comment/:commentID", cc.DeleteComment())
// }
