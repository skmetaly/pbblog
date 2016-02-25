package application

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/config"
	"github.com/skmetaly/pbblog/framework/router"
	"github.com/skmetaly/pbblog/framework/server"
)

type App struct {
	router *httprouter.Router
	server server.Server
	config config.Config
}

func NewApp() *App {
	app := &App{}
	app.SetRouter()
	app.SetServer()
	app.SetConfig()

	return app
}

func (a *App) SetRouter() {
	var r = router.NewRouter()
	a.router = r
}

func (a *App) SetServer() {
	var s = server.NewServer()
	a.server = s
}

func (a *App) SetConfig() {
	//c := NewConfig()
	//a.config = c
}

func (a *App) Run() {
	server.StartServer(a.server, a.router)
}
