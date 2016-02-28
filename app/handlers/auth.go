package handlers

import "net/http"

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	//	Redirect to login if they are not authenticated
	authenticated := false

	if !authenticated {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
	}
}
