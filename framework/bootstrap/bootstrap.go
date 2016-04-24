package bootstrap

import (
	"github.com/skmetaly/pbblog/app/handlers"
	"github.com/skmetaly/pbblog/framework/application"
	"github.com/skmetaly/pbblog/framework/middleware"
	"github.com/skmetaly/pbblog/framework/router"
	"github.com/skmetaly/pbblog/framework/server"

	"net/http"
)

func Run() {
	var app = application.NewApp()

	var adminRouter = router.NewAdminRouter(&app)
	var feRouter = router.NewFERouter(&app)

	middleware := middleware.Middleware{}

	middleware.Add(feRouter)
	middleware.Add(http.HandlerFunc(handlers.AuthenticateRequest))
	middleware.Add(adminRouter)

	server.StartServer(app.Server, middleware)
}
