package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/users"
	"github.com/skmetaly/pbblog/framework/application"
	"github.com/skmetaly/pbblog/framework/validation"
	"net/http"
)

func GETUsersNew(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/users/new", nil)
	}
}

func POSTUsersNew(a application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		user, err := users.NewUser(
			r.FormValue("username"),
			r.FormValue("email"),
			r.FormValue("password"),
		)

		if validation.IsValidationError(err) {
			a.View.Render(w, r, "admin/users/new", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})

			return
		}

		http.Redirect(w, r, "/admin/users?flash=User+created", http.StatusFound)

	}
}
