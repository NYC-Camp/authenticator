// Only contains code for running the service. Brings together logic from various
// other services.
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/nyc-camp/authenticator/libtmpl"
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
	dbmap.AddTableWithName(libuser.User{}, "user").SetKeys(false, "uid")
	userStorageMySQL = UserStorageMySQL{Dbmap: dbmap}
}

func main() {
	loginTmplCfg := libtmpl.HTMLTemplateConfig{TemplateDir: "templates/", DefaultErrorFunc: libuser.HandleError}
	router := httprouter.New()
	recoveryMiddleware := negroni.NewRecovery()
	recoveryMiddleware.PrintStack = false
	n := negroni.New(recoveryMiddleware, negroni.NewLogger(), negroni.NewStatic(http.Dir("public")))

	userRegistration := libuser.UserRegistration{Storage: userStorageMySQL}
	userLogin := libuser.UserLogin{Storage: userStorageMySQL, TemplateConfig: loginTmplCfg}

	router.GET("/register", userRegistration.GetRegistrationForm)
	router.POST("/register", userRegistration.HandleRegistrationSubmission)
	router.GET("/login", userLogin.LoginForm)
	router.POST("/login", userLogin.LoginSubmission)

	n.UseHandler(router)

	n.Run(":4567")
}
