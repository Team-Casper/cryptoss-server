package app

import (
	"encoding/json"
	"net/http"

	"github.com/portto/aptos-go-sdk/models"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/util"
)

type SetProfileReq struct {
	PhoneNumber string         `json:"phone_number"`
	ProfileID   models.TokenID `json:"profile_id"`
}

func (a *App) HandleSetProfile(w http.ResponseWriter, r *http.Request) {
	var reqBody SetProfileReq

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

	account, err := a.GetAccount(reqBody.PhoneNumber)
	if err != nil {
		log.Errorf("failed to get account by the phone number(%s): %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to get account", http.StatusInternalServerError)
		return
	}

	account.ProfileID = reqBody.ProfileID

	if err := a.SetAccount(reqBody.PhoneNumber, account); err != nil {
		log.Errorf("failed to set account by the phone number(%s): %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to get account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
