package application

import (
	// "github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/config"
	"github.com/skmetaly/pbblog/framework/database"
	"github.com/skmetaly/pbblog/framework/server"
	"github.com/skmetaly/pbblog/framework/view"
)

var AppContainerInstance App

type App struct {
	Router   *httprouter.Router
	Server   server.Server
	Config   config.Config
	View     view.View
	Database database.Database
}

func NewApp() App {

	app := &App{}

	app.SetServer()
	app.SetConfig()
	app.SetViews()
	app.SetDatabase()

	AppContainerInstance = *app

	return AppContainerInstance
}

func (a *App) SetServer() {
	var s = server.NewServer()
	a.Server = s
}

func (a *App) SetConfig() {
	a.Config = config.NewConfig()
}

//SetViews sets the view class
func (a *App) SetViews() {
	var v = view.NewView()
	a.View = v
}

//SetDatabase creates a database object and sets the connections
func (a *App) SetDatabase() {
	a.Database = database.Database{
		Config: a.Config.DatabaseConfig,
	}

	a.Database.Connect()
	a.Database.ConnectORM()
}
