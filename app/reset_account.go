package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
)

type ResetAccountReq struct {
	PhoneNumber string `json:"phone_number"`
}

func (a *App) HandleResetAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody ResetAccountReq

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// check validity of request
	if !util.IsValidPhoneNumber(reqBody.PhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.PhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	if err := a.DeleteVerification(reqBody.PhoneNumber); err != nil {
		log.Errorf("failed to delete verification: %v", reqBody.PhoneNumber)
		http.Error(w, "failed to delete verification", http.StatusInternalServerError)
		return
	}

	if err := a.DeleteAccount(reqBody.PhoneNumber); err != nil {
		log.Errorf("failed to delete account: %v", reqBody.PhoneNumber)
		http.Error(w, "failed to delete account", http.StatusInternalServerError)
		return
	}

	if err := a.DeleteDeposit(reqBody.PhoneNumber); err != nil {
		log.Errorf("failed to delete deposit: %v", err)
		http.Error(w, "failed to delete deposit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
