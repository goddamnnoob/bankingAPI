package app

import (
	"log"
	"net/http"

	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
)

func Start() {

	// router/multiplexer to route
	router := mux.NewRouter()

	//wiring together
	uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb())}
	///uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryStub())}

	router.HandleFunc("/getAllUsers", uh.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}", uh.getUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
