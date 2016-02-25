package bootstrap

import (
	"github.com/skmetaly/pbblog/framework/application"
)

func Run() {
	var app = application.NewApp()
	app.Run()
}
