package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skmetaly/pbblog/app/users"
	"github.com/skmetaly/pbblog/framework/application"
	"github.com/skmetaly/pbblog/framework/hash"
	"github.com/skmetaly/pbblog/framework/session"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GETDashboardLogin GET admin/login
func GETDashboardLogin(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		a.View.Render(w, r, "admin/dashboard/login", map[string]interface{}{
			"HideHeader": true,
			"NextURL":    r.URL.Query().Get("next"),
		})
	}
}

// POSTDashboardLogin POST admin/login
func POSTDashboardLogin(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		sessionInstance := session.Instance(r)
		loginAttemptsKey := a.Config.SessionConfig.LoginAttemptsKey
		maxLoginAttempts := a.Config.SessionConfig.MaxLoginAttempts

		// Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid
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

//GETDashboardLogout GET admin/logout
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

// GETDashboardProfile returns the view for the profile of the user
func GETDashboardProfile(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Get session
		sess := session.Instance(r)

		userID := sess.Values["user_id"].(uint)

		uR := &users.UserRepository{Db: a.Database}
		userModel := uR.ByID(userID)

		a.View.Render(w, r, "admin/dashboard/profile", map[string]interface{}{
			"User": userModel,
		})

	}
}

// POSTDashboardProfile handles the POST request for updating the profile of the user
func POSTDashboardProfile(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Get session
		sess := session.Instance(r)

		var err error

		userID := sess.Values["user_id"].(uint)
		uR := &users.UserRepository{Db: a.Database}
		user := uR.ByID(userID)

		// Check if the submission is for personal data change
		if strings.Compare(r.FormValue("action"), "update_personal") == 0 {
			user.Email = r.FormValue("email")
			user.FirstName = r.FormValue("first_name")
			user.LastName = r.FormValue("last_name")

			err = uR.Update(&user)
		}

		// Check if the submission is for password change
		if strings.Compare(r.FormValue("action"), "update_password") == 0 {

			err = users.ValidatePasswordChange(
				uR,
				user,
				r.FormValue("current_password"),
				r.FormValue("new_password"),
				r.FormValue("new_password_verification"))

			if err == nil {
				user.Password = hash.CreateFromPassword(r.FormValue("new_password"))
				err = uR.Update(&user)
			}

		}

		viewParameters := map[string]interface{}{
			"User":    user,
			"Success": "Successfully saved",
			"Title":   "Profile",
		}
		// Check if there were any errors encountered in the update process
		if err != nil {
			viewParameters["Error"] = err.Error()
			delete(viewParameters, "Success")
		}

		a.View.Render(w, r, "admin/dashboard/profile", viewParameters)
	}
}
