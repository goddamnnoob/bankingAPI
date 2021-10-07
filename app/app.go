package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// router/multiplexer to route
	router := mux.NewRouter()

	dbClient := getDbClient()

	//wiring together
	uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb(dbClient))}
	///uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryStub())}
	ah := AccountHandler{service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	router.HandleFunc("/getAllUsers", uh.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", uh.getUsersByStatus).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}", uh.getUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/users/{user_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}

func getDbClient() *sqlx.DB {
	var (
		db_username string = os.Getenv("DATABASE_USERNAME")
		db_password string = os.Getenv("DATABASE_PASSWORD")
		db_name     string = os.Getenv("DATABASE_NAME")
		db_address  string = os.Getenv("DATABASE_SERVER_ADDRESS")
		db_port     string = os.Getenv("DATABASE_SERVER_PORT")
	)
	client, err := sqlx.Open("mysql", db_username+":"+db_password+"@tcp("+db_address+":"+db_port+")/"+db_name)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)
	return client
}
