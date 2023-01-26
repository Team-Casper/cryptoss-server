package handler

import (
	"net/http"
)

func RegisterAddress(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
