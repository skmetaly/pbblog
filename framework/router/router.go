package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app"
	"github.com/skmetaly/pbblog/framework/application"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound

	return router
}

func NewAdminRouter(a application.App) *httprouter.Router {
	router := NewRouter()
	app.AddAdminRoutes(router, a)

	return router
}

func NewFERouter(a application.App) *httprouter.Router {
	router := NewRouter()
	app.AddFERoutes(router, a)

	return router
}
