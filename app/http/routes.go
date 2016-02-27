package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/http/controllers/admin"
	"github.com/skmetaly/pbblog/framework/application"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func AddRoutes(router *httprouter.Router, a application.App) {
	router.GET("/", Index)
	router.GET("/admin/login", admin.GETDashboardLogin(a))
	router.GET("/admin", admin.GETDashboardIndex(a))
}
