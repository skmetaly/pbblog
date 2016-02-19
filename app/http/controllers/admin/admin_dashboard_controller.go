package admin

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GETDashboardIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("test")
}
