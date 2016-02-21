package admin

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"text/template"
)

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
