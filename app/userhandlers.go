package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/errs"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
)

type UserHandlers struct {
	service service.UserService
}

func (uh UserHandlers) getAllUsers(rw http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetAllUsers()

	if err != nil {
		writeResponse(rw, http.StatusInternalServerError, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, users)
	}
}

func (uh UserHandlers) getUsersByStatus(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var users []domain.User
	var err *errs.AppError
	st := 0
	status, b := query["status"]
	if b && status[0] == "active" {
		st = 1
	}
	if !b || len(status) == 0 {
		users, err = uh.service.GetAllUsers()
	} else {
		users, err = uh.service.GetUserByStatus(st)
	}
	if err != nil {
		writeResponse(rw, http.StatusInternalServerError, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, users)
	}
}

func getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
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
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, user)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err)
	}
}
