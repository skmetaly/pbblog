package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/http"
	"github.com/skmetaly/pbblog/framework/application"
)

func NewRouter(a application.App) *httprouter.Router {
	router := httprouter.New()
	http.AddRoutes(router, a)

	return router
}
