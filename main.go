package main

import (
	"TeraApps/config"
	"TeraApps/features/user/data"
	"TeraApps/features/user/handler"
	"TeraApps/features/user/services"
	"TeraApps/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // Koneksi ke DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	//commentData := td.New(db)
	//commentService := ts.NewCommentService(commentData)
	//commentHandler := th.NewHandler(commentService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())         // ini aja cukup
	routes.InitRoute(e, userHandler) //, commentHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
