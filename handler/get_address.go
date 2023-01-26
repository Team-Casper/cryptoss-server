package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const addr = "0xc7ea756470f72ae761b7986e4ed6fd409aad183b1b2d3d2f674d979852f45c4b"

type Address struct {
	Addr string `json:"address"`
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	log.Infof("phone number is %s", mux.Vars(r)["phone-number"])

	address := &Address{
		Addr: addr,
	}

	payload, err := json.Marshal(address)
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
