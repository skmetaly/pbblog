package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/application"
	//"github.com/skmetaly/pbblog/framework/view"
	"net/http"
)

func GETDashboardLogin(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/login", map[string]bool{
			"HideHeader": true,
		})
	}
}

func GETDashboardIndex(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/index", nil)
	}
}
