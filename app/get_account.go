package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
)

type Address struct {
	Addr string `json:"address"`
}

func (a *App) HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	phoneNumber := mux.Vars(r)["phone-number"]

	// check validity of phone number and telecom code
	if !util.IsValidPhoneNumber(phoneNumber) {
		log.Errorf("invalid phone number (%s)", phoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	// get account and set address
	account, err := a.GetAccount(phoneNumber)
	if err != nil {
		log.Errorf("failed to get account by the phone number(%s): %v", phoneNumber, err)
		http.Error(w, "failed to get account by the phone number", http.StatusNotFound)
		return
	}

	payload, err := json.Marshal(account)
	if err != nil {
		log.Errorf("failed to marshal payload")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(payload); err != nil {
		log.Errorf("failed to write response payload")
		return
	}
}
