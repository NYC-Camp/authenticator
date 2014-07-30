// Handles the registration functionality for users
package libuser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
)

// Registration should be simple as we are only collecting a user's username, email address, and password.
// We should ensure the uniqueness checks (for username and email) are fast.
type UserRegistration struct {
	Storage        UserStorage
	TemplateConfig libtmpl.HTMLTemplateConfig
	//@TODO: Should just pass in a session instead of the session store.
	SessionStore sessions.Store
}

type Registration struct {
	Username string `schema:"username"`
	Email    string `schema:"email"'`
	Password string `schema:"password"`
}

func (ur UserRegistration) GetRegistrationForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	htmlTemplateConfig := libtmpl.HTMLTemplateConfig{
		TemplateDir:      "templates/",
		DefaultErrorFunc: ur.HandleError,
	}
	htmlTemplate := htmlTemplateConfig.NewHTMLTemplate()
	htmlTemplate.Content = "templates/registrationform.html"
	htmlTemplate.Execute(w, nil)
}

func (ur UserRegistration) HandleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "An error occured: %v", err)
}

func (ur UserRegistration) HandleRegistrationSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//@TODO: Should just pass in a session instead of the session store.
	userSession, err := ur.SessionStore.Get(r, "authenticator")
	if err != nil {
		log.Printf("session error: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	registration := new(Registration)
	decoder := schema.NewDecoder()

	err = decoder.Decode(registration, r.PostForm)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	// Move this into a convenience  function.
	duplicateUsername, err := ur.Storage.CheckUsername(registration.Username)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	if duplicateUsername {
		//@TODO: Add an invalid username or password flash message once sessions
		//are implemented
		htmlTemplate := ur.TemplateConfig.NewHTMLTemplate()
		htmlTemplate.Content = "templates/registrationform.html"
		htmlTemplate.Execute(w, libtmpl.HTMLTemplateData{
			Content: map[string]interface{}{
				"error":    "This username is already in use.",
				"username": registration.Username,
				"email":    registration.Email,
			},
		})
		return
	}

	validEmail, err := VerifyEmail(registration.Email)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	if !validEmail {
		//@TODO: Add an invalid username or password flash message once sessions
		//are implemented
		htmlTemplate := ur.TemplateConfig.NewHTMLTemplate()
		htmlTemplate.Content = "templates/registrationform.html"
		htmlTemplate.Execute(w, libtmpl.HTMLTemplateData{
			Content: map[string]interface{}{
				"error":    "Email address is not valid.",
				"username": registration.Username,
				"email":    registration.Email,
			},
		})
		return
	}
	// Move this into a convenience  function.
	duplicateEmail, err := ur.Storage.CheckEmail(registration.Email)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	if duplicateEmail {
		//@TODO: Add an invalid username or password flash message once sessions
		//are implemented
		htmlTemplate := ur.TemplateConfig.NewHTMLTemplate()
		htmlTemplate.Content = "templates/registrationform.html"
		htmlTemplate.Execute(w, libtmpl.HTMLTemplateData{
			Content: map[string]interface{}{
				"error":    "This email address is already in use.",
				"username": registration.Username,
				"email":    registration.Email,
			},
		})
		return
	}

	newUser := CreateUser()
	newUser.Email = registration.Email
	newUser.Username = registration.Username
	newUser.SetPassword([]byte(registration.Password))
	_, err = ur.Storage.CreateUser(newUser)
	if err != nil {
		log.Printf("%v\n", err)
	}
	userSession.Values["reg-username"] = registration.Username
	userSession.AddFlash("Your user account has been successfully created. Please check your email for a verification link.")
	userSession.Save(r, w)
	w.Header().Add("Location", "/login")
	w.WriteHeader(http.StatusFound)
	// w.Write([]byte(fmt.Sprintf("<div>username: %v<br>email: %v</div>", newUser.Username, newUser.Email)))

}
