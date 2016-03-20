package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/controllers/admin"
	"github.com/skmetaly/pbblog/framework/application"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func AddFERoutes(router *httprouter.Router, a application.App) {
	router.GET("/", Index)
	router.GET("/admin/login", admin.GETDashboardLogin(a))
}

func AddAdminRoutes(router *httprouter.Router, a application.App) {
	router.GET("/admin", admin.GETDashboardIndex(a))
	router.GET("/admin/users/new", admin.GETUsersNew(a))
	router.POST("/admin/users/new", admin.POSTUsersNew(a))
}
