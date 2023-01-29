package app

import (
	"net/http"
)

func (a *App) HandleSendCoin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
