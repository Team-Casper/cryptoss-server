package app

import (
	"net/http"
)

func (a *App) StartVerification(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
