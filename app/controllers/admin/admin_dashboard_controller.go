package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/users"
	"github.com/skmetaly/pbblog/framework/application"
	"github.com/skmetaly/pbblog/framework/session"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//GETDashboardLogin GET admin/login
func GETDashboardLogin(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/login", map[string]interface{}{
			"HideHeader": true,
			"NextURL":    r.URL.Query().Get("next"),
		})
	}
}

//POSTDashboardLogin POST admin/login
func POSTDashboardLogin(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		sessionInstance := session.Instance(r)
		loginAttemptsKey := a.Config.SessionConfig.LoginAttemptsKey
		maxLoginAttempts := a.Config.SessionConfig.MaxLoginAttempts

		// Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid :-)
		if sessionInstance.Values[loginAttemptsKey] != nil && sessionInstance.Values[loginAttemptsKey].(int) >= maxLoginAttempts {
			log.Println("Brute force login prevented")
			a.View.Render(w, r, "admin/dashboard/index", map[string]interface{}{
				"Error": "Brute force login",
			})

			return
		}

		uR := &users.UserRepository{Db: a.Database}

		success, err := users.LoginUser(sessionInstance, uR, r.FormValue("username"), r.FormValue("password"))

		if err != nil {
			log.Println("Login unsuccessfull" + strconv.FormatBool(success))

			a.View.Render(w, r, "admin/dashboard/login", map[string]interface{}{
				"HideHeader": true,
				"Error":      err.Error(),
				"Username":   r.FormValue("username"),
			})
		}

		//We have a successful login. Save data to session, and redirect user to dashboard
		err = sessionInstance.Save(r, w)
		if err != nil {
			log.Println("Error with saving session")
		}

		//Redirect to "next" if we have a next
		nextURL, err := url.QueryUnescape(r.URL.Query().Get("next"))
		redirectURL := "/admin"

		if nextURL != "" {
			redirectURL = nextURL
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}

//GETDashboardIndex GET admin/
func GETDashboardIndex(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/index", nil)
	}
}

func GETDashboardLogout(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Get session
		sess := session.Instance(r)

		// If user is authenticated
		if sess.Values["user_id"] != nil {
			session.Empty(sess)
			sess.Save(r, w)
		}

		http.Redirect(w, r, "/admin/login", http.StatusFound)

	}

}
