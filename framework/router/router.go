package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app"
	"github.com/skmetaly/pbblog/framework/application"
)

//NewRouter Creates a new httprouter.Router instance
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound

	return router
}

//NewAdminRouter creates a new router with all admin routes added
func NewAdminRouter(a *application.App) *httprouter.Router {
	router := NewRouter()
	app.AddAdminRoutes(router, a)

	return router
}

//NewFERouter creates a new router with all front end routes added
func NewFERouter(a *application.App) *httprouter.Router {
	router := NewRouter()
	app.AddFERoutes(router, a)

	return router
}
