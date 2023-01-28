package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/team-casper/cryptoss-server/config"
	"net/http"
	"time"
)

type App struct {
	DB   *leveldb.DB
	Conf *config.Config
	Srv  *http.Server
}

func New() (*App, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	db, err := leveldb.OpenFile(conf.DBPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error occurs while opening db (%s): %w", conf.DBPath, err)
	}

	return &App{
		Conf: conf,
		DB:   db,
	}, nil
}

func (a *App) InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/verification/start", a.StartVerification).Methods(http.MethodPost)
	router.HandleFunc("/verification/check", a.CheckVerification).Methods(http.MethodPost)
	router.HandleFunc("/address", a.RegisterAddress).Methods(http.MethodPost)
	router.HandleFunc("/address/{phone-number}", a.GetAddress).Methods(http.MethodGet)
	router.HandleFunc("/send-coin", a.SendCoin).Methods(http.MethodPost)
	router.HandleFunc("/profile", a.SetProfile).Methods(http.MethodPost)

	a.Srv = &http.Server{
		Handler:      router,
		Addr:         a.Conf.ListeningAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (a *App) Start() error {
	log.Infof("Cryptoss server is started: %s", a.Conf.ListeningAddr)
	return a.Srv.ListenAndServe()
}

func (a *App) GracefulShutdown() error {
	if err := a.DB.Close(); err != nil {
		log.Errorf("error occurs while closing DB: %v", err)
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return a.Srv.Shutdown(ctxTimeout)
}
