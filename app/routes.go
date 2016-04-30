package app

import (
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/controllers/admin"
	"github.com/skmetaly/pbblog/framework/application"
	"net/http"
)

//  [todo] Move this to config
var adminPrefix = "admin"

//Index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//AddFERoutes Adds all the FE routes that don't need user login
func AddFERoutes(router *httprouter.Router, a *application.App) {
	router.GET("/", Index)
	router.GET("/"+adminPrefix+"/login", admin.GETDashboardLogin(a))
	router.POST("/"+adminPrefix+"/login", admin.POSTDashboardLogin(a))
}

//AddAdminRoutes Adds all the admin routes that need user login
func AddAdminRoutes(router *httprouter.Router, a *application.App) {
	router.GET("/"+adminPrefix, admin.GETDashboardIndex(a))
	router.GET("/"+adminPrefix+"/logout", admin.GETDashboardLogout(a))
	router.GET("/"+adminPrefix+"/users/new", admin.GETUsersNew(a))
	router.POST("/"+adminPrefix+"/users/new", admin.POSTUsersNew(a))
	router.GET("/"+adminPrefix+"/profile", admin.GETDashboardProfile(a))
	router.POST("/"+adminPrefix+"/profile", admin.POSTDashboardProfile(a))
}
