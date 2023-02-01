package app

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/team-casper/cryptoss-server/config"
)

type AptosAcc struct {
	Address string
	Signer  models.SingleSigner
}
type App struct {
	DB        *leveldb.DB
	Conf      *config.Config
	Srv       *http.Server
	AptosCli  client.AptosClient
	EscrowAcc AptosAcc
	ChainID   uint8
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

	aptosCli := client.NewAptosClient(conf.AptosEndpoint)

	escrowSeed, err := hex.DecodeString(strings.TrimPrefix(conf.EscrowSeed, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode escrow seed")
	}

	escrowPrivKey := ed25519.NewKeyFromSeed(escrowSeed)
	escrowSigner := models.NewSingleSigner(escrowPrivKey)

	chainIDUint64, err := strconv.ParseUint(conf.AptosChainID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse chain ID")
	}

	return &App{
		Conf:     conf,
		DB:       db,
		AptosCli: aptosCli,
		EscrowAcc: AptosAcc{
			Address: hex.EncodeToString(escrowSigner.AccountAddress[:]),
			Signer:  escrowSigner,
		},
		ChainID: uint8(chainIDUint64),
	}, nil
}

func (a *App) InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/verification/start", a.HandleStartVerification).Methods(http.MethodPost)
	router.HandleFunc("/verification/check", a.HandleCheckVerification).Methods(http.MethodPost)
	router.HandleFunc("/account/address", a.HandleRegisterAddress).Methods(http.MethodPost)
	router.HandleFunc("/account/{phone-number}", a.HandleGetAccount).Methods(http.MethodGet)
	router.HandleFunc("/escrow", a.HandleSendToEscrow).Methods(http.MethodPost)
	router.HandleFunc("/profile", a.HandleSetProfile).Methods(http.MethodPost)
	router.HandleFunc("/reset", a.HandleResetAccount).Methods(http.MethodPost)
	router.HandleFunc("/profile", a.HandleSetProfile).Methods(http.MethodPost)

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
