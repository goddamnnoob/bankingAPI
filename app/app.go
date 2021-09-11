package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/getAllUsers", getAllUsers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
