package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/http/controllers/admin"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func AddRoutes(router *httprouter.Router) {
	router.GET("/", Index)
	router.GET("/admin/login", admin.GETDashboardLogin)
	router.GET("/admin", admin.GETDashboardIndex)
}
