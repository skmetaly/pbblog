package handlers

import (
	"github.com/skmetaly/pbblog/framework/session"
	"net/http"
)

//AuthenticateRequest checks if for a given requrest the user is authenticated or not
func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	//Redirect to login if they are not authenticated
	//Get session
	sess := session.Instance(r)

	//If user is not authenticated, don't allow them to access the page
	if sess.Values["user_id"] == nil {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
	}
}
