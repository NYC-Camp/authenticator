// Only contains code for running the service. Brings together logic from various
// other services.
package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Authenticator!"))
	})

	n.UseHandler(router)

	n.Run(":4567")
}
