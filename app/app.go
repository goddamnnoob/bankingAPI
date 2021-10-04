package app

import (
	"log"
	"net/http"
	"os"

	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
)

func Start() {
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// router/multiplexer to route
	router := mux.NewRouter()

	//wiring together
	uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb())}
	///uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryStub())}

	router.HandleFunc("/getAllUsers", uh.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", uh.getUsersByStatus).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}", uh.getUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}
