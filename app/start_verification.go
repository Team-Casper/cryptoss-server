package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/sms"
	"github.com/team-casper/cryptoss-server/types"
	"github.com/team-casper/cryptoss-server/util"
	"net/http"
	"time"
)

type VerificationStartReq struct {
	Nickname    string `json:"nickname"`
	PhoneNumber string `json:"phone_number"`
	TelecoCode  string `json:"teleco_code"`
}

type VerificationStartResp struct {
	VerificationCode string `json:"verification_code"`
}

func (a *App) HandleStartVerification(w http.ResponseWriter, r *http.Request) {
	// get request body
	var reqBody VerificationStartReq

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// check validity of phone number and telecom code
	if !util.IsValidPhoneNumber(reqBody.PhoneNumber) {
		log.Errorf("invalid phone number (%s)", reqBody.PhoneNumber)
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	if !util.IsValidTelecoCode(reqBody.TelecoCode) {
		log.Errorf("invalid telecom code (%s)", reqBody.TelecoCode)
		http.Error(w, "invalid telecom code", http.StatusBadRequest)
		return
	}

	// generate verification code(random 6-digit number)
	code, err := util.GenVerificationCode(6)
	if err != nil {
		log.Errorf("error occurs when generating verification code: %v", err)
		http.Error(w, "error occurs when generating verification code", http.StatusInternalServerError)
		return
	}

	// set verification code
	newVerification := reqBody.toVerification(code)
	if err := a.SetVerification(reqBody.PhoneNumber, newVerification); err != nil {
		log.Errorf("failed to set verification: %v", err)
		http.Error(w, "failed to set verification", http.StatusInternalServerError)
		return
	}

	// send sms
	if err := sms.Send(a.Conf, reqBody.PhoneNumber, code); err != nil {
		log.Errorf("failed to send sms to %s: %v", reqBody.PhoneNumber, err)
		http.Error(w, "failed to send sms", http.StatusInternalServerError)
		return
	}

	// response
	payload, err := json.Marshal(VerificationStartResp{VerificationCode: code})
	if err != nil {
		log.Errorf("failed to marshal response: %v", err)
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(payload); err != nil {
		log.Errorf("failed to write response payload: %s", err.Error())
		return
	}
}

func (r *VerificationStartReq) toVerification(code string) *types.Verification {
	return &types.Verification{
		Nickname:   r.Nickname,
		TelecoCode: r.TelecoCode,
		Code:       code,
		Expiry:     time.Now().Add(time.Minute * 5),
	}
}
