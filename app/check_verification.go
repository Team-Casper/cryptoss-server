package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/types"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
	"strings"
)

type VerificationCheckReq struct {
	PhoneNumber      string `json:"phone_number"`
	VerificationCode string `json:"verification_code"`
}

func (a *App) CheckVerification(w http.ResponseWriter, r *http.Request) {
	// get request body
	var reqBody VerificationCheckReq

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validate phone number
	if !util.IsValidPhoneNumber(reqBody.PhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.PhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	// get verification code
	verification, err := a.GetVerification(reqBody.PhoneNumber)
	if err != nil {
		log.Errorf("failed to get verification code under the number of %s: %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to get verification code", http.StatusBadRequest)
		return
	}

	if strings.Compare(verification.Code, reqBody.VerificationCode) != 0 {
		log.Errorf("invalid verification code, expected (%s) but got (%s)", verification.Code, reqBody.VerificationCode)
		http.Error(w, "invalid verification code", http.StatusUnauthorized)
		return
	}

	// check expiry
	if verification.IsExpired() {
		log.Errorf("verification code expired")
		http.Error(w, "verification code expired", http.StatusUnauthorized)
		return
	}

	// set account
	newAccount := types.NewAccount(verification.Nickname, "")
	if err := a.SetAccount(reqBody.PhoneNumber, newAccount); err != nil {
		log.Errorf("failed to set account by the phone number of %s: %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to set account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
