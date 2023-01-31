package app

import (
	"context"
	"fmt"
	"github.com/portto/aptos-go-sdk/crypto"
	"github.com/portto/aptos-go-sdk/models"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var addr0x1 models.AccountAddress
var aptosCoinTypeTag models.TypeTag

func init() {
	addr0x1, _ = models.HexToAccountAddress("0x1")
	aptosCoinTypeTag = models.TypeTagStruct{
		Address: addr0x1,
		Module:  "aptos_coin",
		Name:    "AptosCoin",
	}
}

func (a *App) TransferCoin(receiverAddress string, receiverPubKey []byte, amount uint64) error {
	ctx := context.Background()

	// check if account exists
	existing, err := a.isExistingAccount(ctx, receiverAddress)
	if err != nil {
		return fmt.Errorf("error occurs while checking account exists")
	}

	// if not, create a new one
	if !existing {
		// create account
		if err := a.createAccountTx(ctx, receiverPubKey); err != nil {
			return fmt.Errorf("failed to create account: %w", err)
		}
	}

	// send coin
	if err := a.sendCoinTx(ctx, receiverAddress, amount); err != nil {
		return fmt.Errorf("failed to send APT: %w", err)
	}

	return nil
}

func getTransferPayload(to models.AccountAddress, amount uint64) models.TransactionPayload {
	return models.EntryFunctionPayload{
		Module: models.Module{
			Address: addr0x1,
			Name:    "coin",
		},
		Function:      "transfer",
		TypeArguments: []models.TypeTag{aptosCoinTypeTag},
		Arguments:     []interface{}{to, amount},
	}
}

func (a *App) isExistingAccount(ctx context.Context, address string) (bool, error) {
	if _, err := a.AptosCli.GetAccount(ctx, address); err != nil {
		if strings.Contains(err.Error(), "account_not_found") {
			return false, nil
		}
		return false, fmt.Errorf("failed to get account: %w", err)
	}
	return true, nil
}

func (a *App) createAccountTx(ctx context.Context, targetPubKey []byte) error {
	authKey := crypto.SingleSignerAuthKey(targetPubKey)

	seqNum, err := a.getEscrowSeqNum(ctx)
	if err != nil {
		return fmt.Errorf("failed to get sequence number of escrow account: %w", err)
	}

	tx := models.Transaction{}

	err = tx.SetChainID(a.ChainID).
		SetSender(a.EscrowAcc.Address).
		SetPayload(models.EntryFunctionPayload{
			Module: models.Module{
				Address: addr0x1,
				Name:    "aptos_account",
			},
			Function:  "create_account",
			Arguments: []interface{}{authKey},
		}).SetExpirationTimestampSecs(uint64(time.Now().Add(10 * time.Minute).Unix())).
		SetGasUnitPrice(uint64(100)).
		SetMaxGasAmount(uint64(5000)).
		SetSequenceNumber(seqNum).Error()
	if err != nil {
		return fmt.Errorf("failed to build tx to create account: %w", err)
	}

	if err := a.EscrowAcc.Signer.Sign(&tx).Error(); err != nil {
		return fmt.Errorf("failed to sign tx message: %w", err)
	}

	rawTx, err := a.AptosCli.SubmitTransaction(ctx, tx.UserTransaction)
	if err != nil {
		panic(err)
	}

	log.Infof("created account (%s):", rawTx.Hash)

	return nil
}

func (a *App) sendCoinTx(ctx context.Context, receiverAddress string, amount uint64) error {
	receiverAccount, err := models.HexToAccountAddress(receiverAddress)
	if err != nil {
		return fmt.Errorf("failed to convert to hex address: %w", err)
	}

	seqNum, err := a.getEscrowSeqNum(ctx)
	if err != nil {
		return fmt.Errorf("failed to get sequence number of escrow account: %w", err)
	}

	txTransfer := models.Transaction{}
	err = txTransfer.SetChainID(a.ChainID).
		SetSender(a.EscrowAcc.Address).
		SetPayload(getTransferPayload(receiverAccount, amount)).
		SetExpirationTimestampSecs(uint64(time.Now().Add(10 * time.Minute).Unix())).
		SetGasUnitPrice(uint64(100)).
		SetMaxGasAmount(uint64(5000)).
		SetSequenceNumber(seqNum).Error()
	if err != nil {
		return fmt.Errorf("failed to build tx: %w", err)
	}

	if err := a.EscrowAcc.Signer.Sign(&txTransfer).Error(); err != nil {
		return fmt.Errorf("failed to sign tx message: %w", err)
	}

	rawTxTransfer, err := a.AptosCli.SubmitTransaction(ctx, txTransfer.UserTransaction)
	if err != nil {
		return fmt.Errorf("failed to submit tx: %w", err)
	}

	log.Infof("coin transferred. tx hash(%s)", rawTxTransfer.Hash)

	return nil
}

func (a *App) getEscrowSeqNum(ctx context.Context) (string, error) {
	escrowAccount, err := a.AptosCli.GetAccount(ctx, a.EscrowAcc.Address)
	if err != nil {
		return "", fmt.Errorf("failed to get escrow account: %w", err)
	}

	return escrowAccount.SequenceNumber, nil
}
