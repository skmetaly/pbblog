package bootstrap

import (
	"github.com/skmetaly/pbblog/framework/application"
	"github.com/skmetaly/pbblog/framework/router"
	"github.com/skmetaly/pbblog/framework/server"
)

func Run() {
	var app = application.NewApp()

	var r = router.NewRouter(app)
	server.StartServer(app.Server, r)
}
