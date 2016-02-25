package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/http"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	http.AddRoutes(router)

	return router
}
