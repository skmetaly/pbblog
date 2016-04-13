package session

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	Store *sessions.CookieStore //Store is the cookie store
	Name  string                //Name is the session name
)

//SessionConfig stores session level information
type SessionConfig struct {
	Options          sessions.Options `json:"Options"`          // Pulled from: http://www.gorillatoolkit.org/pkg/sessions#Options
	Name             string           `json:"Name"`             // Name for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.Get
	SecretKey        string           `json:"SecretKey"`        //
	MaxLoginAttempts int              `json:"MaxLoginAttempts"` //
	LoginAttemptsKey string           `json:"LoginAttemptsKey"` //
}

//Configure the session cookie store
func Configure(s SessionConfig) {
	Store = sessions.NewCookieStore([]byte(s.SecretKey))
	Store.Options = &s.Options
	Name = s.Name
}

//Instance returns a new session, never returns an error
func Instance(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, Name)
	return session
}

//Empty deletes all the current session values
func Empty(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}
