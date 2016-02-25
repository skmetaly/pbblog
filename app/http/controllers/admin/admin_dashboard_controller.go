package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/framework/view"
	"net/http"
	"os"
	"text/template"
)

//	GET admin/
func GETDashboardIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	responseTemplate, err := template.New("DashboardIndex").Parse("<h1>Yey with {{.}}!</h1>")
	if err != nil {
		panic(err)
	}

	err = responseTemplate.Execute(os.Stdout, "AAAGETDashboardIndex")
	if err != nil {
		panic(err)
	}
}

//	GET admin/login
func GETDashboardLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	view.Render(w, "admin/dashboard/login.html")
}
