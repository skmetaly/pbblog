package framework

import (
	"github.com/julienschmidt/httprouter"
)

type App struct {
	router *httprouter.Router
	server Server
}

func NewApp() *App {
	app := &App{}
	app.SetRouter()
	app.SetServer()

	return app
}

func (a *App) SetRouter() {
	var r = NewRouter()
	a.router = r
}

func (a *App) SetServer() {
	var s = NewServer()
	a.server = s
}

func (a *App) Run() {
	StartServer(a.server, a.router)
}
