package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name" xml:"name"`
	UID      int64  `json:"uid" xml:"uid"`
	Username string `json:"username" xml:"username"`
}

type UserHandlers struct {
	service service.UserService
}

func (uh UserHandlers) getAllUsers(rw http.ResponseWriter, r *http.Request) {
	users, _ := uh.service.GetAllUsers()

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

func (uh *UserHandlers) getUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]
	user, err := uh.service.GetUser(id)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, err.Error())
	} else {
		rw.Header().Add("ContenT-Type", "application/json")
		json.NewEncoder(rw).Encode(user)
	}
}
