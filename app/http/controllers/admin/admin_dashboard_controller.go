package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/application"
	//"github.com/skmetaly/pbblog/framework/view"
	"net/http"
)

func GETDashboardLogin(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/logindd", nil)

	}
}

func GETDashboardIndex(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/login", nil)

	}
}
