package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CheckVerification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Errorf("method of request is invalid")
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
