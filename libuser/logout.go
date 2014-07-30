// Handles the user loggout functionality
package libuser

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
)

type UserLogout struct {
	Storage        UserStorage
	TemplateConfig libtmpl.HTMLTemplateConfig
	//@TODO: Should just pass in a session instead of the session store.
	SessionStore sessions.Store
}

func (ul UserLogout) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//@TODO: Should just pass in a session instead of the session store.
	userSession, err := ul.SessionStore.Get(r, "authenticator")
	if err != nil {
		log.Printf("session error: %v", err)
	}
	if userSession.Values["logged-in"] == true {
		userSession.Values = make(map[interface{}]interface{})
		userSession.AddFlash("Successfully Logged Out")
		userSession.Save(r, w)
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
}
