// Only contains code for running the service. Brings together logic from various
// other services.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nyc-camp/authenticator/libuser"
)

var userStorageMySQL UserStorageMySQL

func init() {
	db, err := sql.Open("mysql", "authenticator:authenticator@tcp(localhost:3306)/authenticator?parseTime=true")
	if err != nil {
		log.Printf("%v\n", err)
		log.Panic("An error occured while connecting to the database.")
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	table := dbmap.AddTableWithName(libuser.User{}, "user").SetKeys(false, "uid")
	log.Printf("%v\n", table)
	userStorageMySQL = UserStorageMySQL{Dbmap: dbmap}
}

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		aUser := libuser.CreateUser()
		aUser.Email = "hello@world.com"
		aUser.Username = "hello"
		aUser.SetPassword([]byte("Hello"))
		success, err := userStorageMySQL.CreateUser(aUser)
		if err != nil {
			log.Printf("%v\n", err)
		}

		w.Write([]byte("Welcome to the Authenticator!"))
		w.Write([]byte(fmt.Sprintf("<div>%v</div>", success)))
	})

	n.UseHandler(router)

	n.Run(":4567")
}
