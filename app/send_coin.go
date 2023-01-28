package app

import (
	"net/http"
)

func (a *App) SendCoin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
