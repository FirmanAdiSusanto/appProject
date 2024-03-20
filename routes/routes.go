package routes

import (
	"21-api/config"
	comment "21-api/features/comment"
	user "21-api/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, cc comment.CommentController) {
	userRoute(c, ctl)
	//activityRoute(c, ac)

}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/login", ctl.Login())
	c.GET("/users/:hp", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

	//Register User
	c.POST("/register", ctl.RegisterUser()) //Endpoint untuk API

	//DeleteUser
	c.DELETE("/users/:id", ctl.DeleteUser(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

}

func commentRoute(c *echo.Echo, cc comment.CommentController) {
	//Komentar
	c.POST("/comment", cc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
