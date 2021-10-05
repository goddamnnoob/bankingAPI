package app

import (
	"encoding/json"
	"net/http"

	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(rw http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		request.UserId = user_id
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message)
		} else {
			writeResponse(rw, http.StatusCreated, account)
		}
	}
}
