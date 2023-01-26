package handler

import (
	"net/http"
)

func SetProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
