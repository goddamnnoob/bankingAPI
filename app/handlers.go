package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name" xml:"name"`
	UID      int64  `json:"uid" xml:"uid"`
	Username string `json:"username" xml:"username"`
}

func getAllUsers(rw http.ResponseWriter, r *http.Request) {
	users := []User{{"Gow", 1, "lol"}, {"ldf", 2, "aaaaa"}, {"aaaa", 3, "wwwwwwwww"}}

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(users)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(users)
	}

}

func getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Contebt-Type", "application/json")
	json.NewEncoder(rw).Encode(mux.Vars(r)["uid"])
}

func createUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "User Created")
}
