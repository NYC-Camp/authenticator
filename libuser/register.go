// Handles the registration functionality for users
package libuser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
)

// Registration should be simple as we are only collecting a user's username, email address, and password.
// We should ensure the uniqueness checks (for username and email) are fast.
type UserRegistration struct {
	Storage UserStorage
}

type Registration struct {
	Username string `schema:"username"`
	Email    string `schema:"email"'`
	Password string `schema:"password"`
}

func (ur UserRegistration) GetRegistrationForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	htmlTemplateConfig := libtmpl.HTMLTemplateConfig{TemplateDir: "templates/", DefaultErrorFunc: ur.HandleError}
	htmlTemplate := htmlTemplateConfig.NewHTMLTemplate()
	htmlTemplate.Content = "templates/registrationform.html"
	htmlTemplate.Execute(w, nil)
}

func (ur UserRegistration) HandleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "An error occured: %v", err)
}

func (ur UserRegistration) HandleRegistrationSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
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

	log.Printf("Username: %v", duplicateUsername)
	if duplicateUsername {
		w.Write([]byte("Username already in use."))
		return
	}

	validEmail, err := VerifyEmail(registration.Username)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	if !validEmail {
		w.Write([]byte("Email not valid."))
		return
	}
	// Move this into a convenience  function.
	duplicateEmail, err := ur.Storage.CheckEmail(registration.Email)
	if err != nil {
		ur.HandleError(w, err)
		return
	}

	if duplicateEmail {
		w.Write([]byte("Email already in use."))
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

	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("Welcome to the Authenticator!"))
	w.Write([]byte(fmt.Sprintf("<div>username: %v<br>email: %v</div>", newUser.Username, newUser.Email)))

}
