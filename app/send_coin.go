package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/types"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
	"strconv"
)

type SendToEscrowReq struct {
	SenderPhoneNumber   string `json:"sender_phone_number"`
	ReceiverPhoneNumber string `json:"receiver_phone_number"`
	Amount              string `json:"amount"`
}

func (a *App) HandleSendToEscrow(w http.ResponseWriter, r *http.Request) {
	// get request body
	var reqBody SendToEscrowReq

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// check validity of request
	if !util.IsValidPhoneNumber(reqBody.SenderPhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.SenderPhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	if !util.IsValidPhoneNumber(reqBody.ReceiverPhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.ReceiverPhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseUint(reqBody.Amount, 10, 64)
	if err != nil {
		log.Errorf("failed to parse amount to int(%s): %v", reqBody.Amount, err)
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	// TODO: check deposit by querying tx by hash

	// set deposit store
	newDeposit := types.NewDeposit(reqBody.SenderPhoneNumber, amount)

	if err := a.SetDeposit(reqBody.ReceiverPhoneNumber, newDeposit); err != nil {
		log.Errorf("failed to set deposit under the phone number(%s): %v", reqBody.ReceiverPhoneNumber, err)
		http.Error(w, "failed to set deposit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
