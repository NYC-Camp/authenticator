// Handles the login functionality for users.
package libuser

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
)

// Login functionality. Should handle checking the user's username and password. Should be able to
// easily take a username or email address to improve the user experience.
type UserLogin struct {
	Storage        UserStorage
	TemplateConfig libtmpl.HTMLTemplateConfig
	//@TODO: Should just pass in a session instead of the session store.
	SessionStore sessions.Store
}

type Login struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}

func (ul UserLogin) LoginForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//@TODO: Should just pass in a session instead of the session store.
	userSession, err := ul.SessionStore.Get(r, "authenticator")
	if err != nil {
		log.Printf("session error: %v", err)
	}
	flashes := userSession.Flashes()
	username := userSession.Values["reg-username"]
	if username != "" {
		delete(userSession.Values, "reg-username")
	}
	userSession.Save(r, w)

	htmlTemplate := ul.TemplateConfig.NewHTMLTemplate()
	htmlTemplate.Content = "templates/loginform.html"
	htmlTemplate.Execute(w, libtmpl.HTMLTemplateData{
		Content: map[string]interface{}{
			"flashes":  flashes,
			"username": username,
		},
	})
}

func (ul UserLogin) HandleError(w http.ResponseWriter, err error) {
	HandleError(w, err)
}

func (ul UserLogin) LoginSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//@TODO: Should just pass in a session instead of the session store.
	userSession, err := ul.SessionStore.Get(r, "authenticator")
	if err != nil {
		log.Printf("session error: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		ul.HandleError(w, err)
		return
	}

	login := new(Login)
	decoder := schema.NewDecoder()

	err = decoder.Decode(login, r.PostForm)
	if err != nil {
		ul.HandleError(w, err)
		return
	}

	// Attempt Password Verficiation
	success, err := ul.Storage.CheckPassword(login.Username, login.Password)
	// Return success or error
	if err != nil {
		//@TODO: Add an invalid username or password flash message once sessions
		//are implemented
		htmlTemplate := ul.TemplateConfig.NewHTMLTemplate()
		htmlTemplate.Content = "templates/loginform.html"
		htmlTemplate.Execute(w, libtmpl.HTMLTemplateData{
			Content: map[string]interface{}{
				"error":    "The username or password you used is not valid.",
				"username": login.Username,
			},
		})
		return
	}

	if success {
		user, err := ul.Storage.RetrieveUser(login.Username)
		if err != nil {
			ul.HandleError(w, err)
			return
		}
		user.LastLogin = time.Now()
		_, err = ul.Storage.UpdateUser(user)
		if err != nil {
			log.Printf("%v", err)
		}
		userSession.Values["logged-in"] = true
		userSession.Values["user"] = user
		userSession.Save(r, w)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte("Welcome to the Authenticator!"))
		w.Write([]byte(fmt.Sprintf("<div>username: %v<br>email: %v</div>", user.Username, user.Email)))
	}
}
