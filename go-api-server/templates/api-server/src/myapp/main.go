package main

import (
	"log"
	"net/http"
	"os"

	"utils"

	"github.com/gorilla/mux"
)

//App application to run the server
type App struct {
	router *mux.Router
}

// main
// entry point of the program
func main() {
	// define the URL handler
	a := App{}
	a.Init()
	// launch the http server
	myport := os.Getenv("PORT")
	if myport == "" {
		myport = ":8080"
	} else {
		myport = ":" + myport
	}
	log.Printf("Server started on %s \n", myport)
	log.Fatal(http.ListenAndServe(myport, a.router))
}

// Init init the http server
func (a *App) Init() {
	a.router = mux.NewRouter().StrictSlash(true)
	a.router.HandleFunc("/", index).Methods(http.MethodGet)
	a.router.HandleFunc("/health", health).Methods(http.MethodGet)
	a.router.HandleFunc("/liveness", liveness).Methods(http.MethodGet)
	initHandlers(a.router)
	a.router.Use(mux.CORSMethodMiddleware(a.router))
}

// initHandlers initialises the URL handlers
// router the router hosting the handlers
// dbconn the DB connection used to interact with database
func initHandlers(router *mux.Router) {
	//sub := router.PathPrefix("/api/v1").Subrouter()

}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func index(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "Hello from go api server"}
	utils.RespondJSON(w, r, 200, msg)
}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func liveness(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"status": "UP"}
	utils.RespondJSON(w, r, 200, msg)
}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func health(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"health": "OK"}
	utils.RespondJSON(w, r, 200, msg)
}
