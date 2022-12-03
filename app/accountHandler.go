package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jimmmisss/banking/dto"
	"github.com/jimmmisss/banking/service"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		account, errApp := a.service.MakeTransaction(request)

		if errApp != nil {
			writeResponse(w, errApp.Code, errApp.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}

func (a AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := a.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}

	}
}
