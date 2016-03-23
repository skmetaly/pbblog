package application

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/config"
	"github.com/skmetaly/pbblog/framework/server"
	"github.com/skmetaly/pbblog/framework/view"
)

type App struct {
	Router *httprouter.Router
	Server server.Server
	Config config.Config
	View   view.View
}

func NewApp() App {
	app := &App{}

	app.SetServer()
	app.SetConfig()
	app.SetViews()

	return *app
}

func (a *App) SetServer() {
	var s = server.NewServer()
	a.Server = s
}

func (a *App) SetConfig() {
	a.Config = config.NewConfig()
}

func (a *App) SetViews() {
	var v = view.NewView()
	a.View = v
}
