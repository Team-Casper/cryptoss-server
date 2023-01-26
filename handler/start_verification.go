package handler

import (
	"net/http"
)

func StartVerification(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
