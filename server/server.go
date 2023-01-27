package server

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/handler"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func New() *Server {
	router := mux.NewRouter()

	router.HandleFunc("/verification/start", handler.StartVerification).Methods(http.MethodPost)
	router.HandleFunc("/verification/check", handler.CheckVerification).Methods(http.MethodGet)
	router.HandleFunc("/address", handler.RegisterAddress).Methods(http.MethodPost)
	router.HandleFunc("/address/{phone-number}", handler.GetAddress).Methods(http.MethodGet)
	router.HandleFunc("/send-coin", handler.SendCoin).Methods(http.MethodPost)
	router.HandleFunc("/profile", handler.SetProfile).Methods(http.MethodPost)

	return &Server{
		&http.Server{
			Handler:      router,
			Addr:         "0.0.0.0:8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (srv *Server) Run() error {
	log.Infof("Cryptoss server is started: %s", srv.Addr)
	return srv.ListenAndServe()
}
