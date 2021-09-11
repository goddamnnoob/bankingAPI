package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// router/multiplexer to route
	router := mux.NewRouter()

	router.HandleFunc("/getAllUsers", getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{uid:[0-9]+}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/createUser", createUser).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
