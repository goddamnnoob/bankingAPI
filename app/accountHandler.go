package app

import (
	"encoding/json"
	"net/http"

	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/logger"
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

func (ah AccountHandler) MakeTransaction(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	userId := vars["user_id"]
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	logger.Error("Transaction type" + request.TransactionType)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.UserId = userId

		account, appError := ah.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.AsMessage())
		} else {
			writeResponse(rw, http.StatusOK, account)
		}
	}

}
