package app

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/util"
)

type RegisterAddressReq struct {
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	PubKeyHex   string `json:"pub_key_hex"`
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

	if len(reqBody.PubKeyHex) == 0 {
		log.Errorf("public key is empty")
		http.Error(w, "public key is empty", http.StatusBadRequest)
		return
	}

	// TODO: check validity of address
	account.Address = reqBody.Address
	account.PubKeyHex = reqBody.PubKeyHex

	if err := a.SetAccount(reqBody.PhoneNumber, account); err != nil {
		log.Errorf("failed to set account: %v", err)
		http.Error(w, "failed to set account", http.StatusInternalServerError)
		return
	}

	// if deposit exists, send asset to the account
	if a.HasClaimableDeposit(reqBody.PhoneNumber) {
		// send deposit to account
		deposit, err := a.GetDeposit(reqBody.PhoneNumber)
		if err != nil {
			log.Errorf("failed to get deposit under %s: %v", reqBody.PhoneNumber, err)
			http.Error(w, "failed to get deposit", http.StatusInternalServerError)
			return
		}

		pubKeyStr := strings.TrimPrefix(reqBody.PubKeyHex, "0x")
		receiverPubKey, err := hex.DecodeString(pubKeyStr)
		if err != nil {
			log.Errorf("failed to decode pub key(%s): %v", reqBody.PubKeyHex, err)
			http.Error(w, "failed to decode pub key", http.StatusInternalServerError)
			return
		}

		if err := a.TransferCoin(reqBody.PhoneNumber, receiverPubKey, deposit.Amount); err != nil {
			log.Errorf("failed to send deposit: %v", err)
			http.Error(w, "failed to send deposit", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
