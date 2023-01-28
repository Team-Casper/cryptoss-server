package app

import (
	"net/http"
)

func (a *App) RegisterAddress(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
