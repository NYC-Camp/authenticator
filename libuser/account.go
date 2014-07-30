// Handles the user account page
package libuser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
)

type UserAccount struct {
	Storage        UserStorage
	TemplateConfig libtmpl.HTMLTemplateConfig
	//@TODO: Should just pass in a session instead of the session store.
	SessionStore sessions.Store
}

func (ua UserAccount) AccountPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//@TODO: Should just pass in a session instead of the session store.
	userSession, err := ua.SessionStore.Get(r, "authenticator")
	if err != nil {
		log.Printf("session error: %v", err)
	}
	if userSession.Values["logged-in"] != true {
		log.Printf("logged-in: %v", userSession.Values)
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	user, ok := userSession.Values["user"].(*User)
	if !ok {
		log.Printf("user: %v", userSession.Values["user"])
		delete(userSession.Values, "logged-in")
		userSession.AddFlash("Please log in")
		userSession.Save(r, w)
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("<a href=\"/logout\">Logout</a>"))
	w.Write([]byte("<h1>Welcome to the Authenticator!</h1>"))
	w.Write([]byte(fmt.Sprintf("<div>username: %v<br>email: %v</div>", user.Username, user.Email)))
}
