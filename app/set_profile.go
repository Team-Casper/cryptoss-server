package app

import (
	"net/http"
)

func (a *App) SetProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
