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

func GETDashboardProfile(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Get session
		sess := session.Instance(r)

		if sess.Values["user_id"] != nil {
			userID := sess.Values["user_id"].(uint)
			/*
				if err != nil {
					http.Redirect(w, r, "/admin/login", http.StatusFound)
					return
				}
			*/
			uR := &users.UserRepository{Db: a.Database}
			userModel := uR.ByID(userID)

			a.View.Render(w, r, "admin/dashboard/profile", map[string]interface{}{
				"User": userModel,
			})

		}
	}
}

func POSTDashboardProfile(a *application.App) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Get session
		sess := session.Instance(r)

		if sess.Values["user_id"] != nil {
			userID := sess.Values["user_id"].(uint)

			uR := &users.UserRepository{Db: a.Database}
			user := uR.ByID(userID)

			//Check if the submission is for email and not for password
			if r.FormValue("email") != "" {
				user.Email = r.FormValue("email")
				uR.Update(user)
			}

			if r.FormValue("current_password") != "" {

				currentPassword := r.FormValue("current_password")

				//Check if passwords are matching
				passMatch := hash.CompareWithHash([]byte(user.Password), currentPassword)

				if passMatch == false {
					a.View.Render(w, r, "admin/dashboard/profile", map[string]interface{}{
						"User":  user,
						"Error": "Password doesn't match",
						"Title": "Profile",
					})
				} else {

					//Current password is true, update the existing password
					user.Password = hash.CreateFromPassword(r.FormValue("new_password"))
					uR.Update(user)

				}

			}

			a.View.Render(w, r, "admin/dashboard/profile", map[string]interface{}{
				"User":    user,
				"Success": "Successfully saved",
				"Title":   "Profile",
			})

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
	}
}
