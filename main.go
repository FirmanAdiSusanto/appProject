package main

import (
	"21-api/config"
	cr "21-api/features/comment/data"
	ch "21-api/features/comment/handler"
	cs "21-api/features/comment/service"
	pr "21-api/features/post/data"
	ph "21-api/features/post/handler"
	ps "21-api/features/post/service"
	ur "21-api/features/user/data"
	uh "21-api/features/user/handler"
	us "21-api/features/user/service"
	"21-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	uq := ur.New(db) // bagian yang menghungkan coding kita ke database / bagian dimana kita ngoding untk ke DB
	us := us.NewService(uq)
	uh := uh.NewUserHandler(us)

	cq := cr.New(db) // bagian yang menghungkan coding kita ke database / bagian dimana kita ngoding untk ke DB
	cs := cs.NewCommentService(cq)
	ch := ch.NewHandler(cs)
	// bagian yang menghandle segala hal yang berurusan dengan HTTP / echo

	pq := pr.New(db)
	ps := ps.NewPostService(pq)
	ph := ph.NewHandler(ps)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, uh, ph, ch)
	e.Logger.Fatal(e.Start(":1323"))
}
