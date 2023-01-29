package app

import (
	"net/http"
)

func (a *App) HandleSetProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
