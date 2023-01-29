package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
)

type RegisterAddressReq struct {
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (a *App) HandleRegisterAddress(w http.ResponseWriter, r *http.Request) {
	// get request body
	var reqBody RegisterAddressReq

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// check validity of phone number
	if !util.IsValidPhoneNumber(reqBody.PhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.PhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	// get account and set address
	account, err := a.GetAccount(reqBody.PhoneNumber)
	if err != nil {
		log.Errorf("failed to get account by the phone number(%s): %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to get account by the phone number", http.StatusNotFound)
		return
	}

	// TODO: check validity of address
	account.Address = reqBody.Address

	if err := a.SetAccount(reqBody.PhoneNumber, account); err != nil {
		log.Errorf("failed to set account: %v", err)
		http.Error(w, "failed to set account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
