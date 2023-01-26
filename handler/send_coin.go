package handler

import (
	"net/http"
)

func SendCoin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
