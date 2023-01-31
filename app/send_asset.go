package app

import (
	"fmt"

	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"
	"github.com/team-casper/cryptoss-server/types"
)

const (
	TestnetChainID = 2
)

var addr0x1 models.AccountAddress
var aptosCoinTypeTag models.TypeTag
var escrowAddress string
var escrowAddr models.AccountAddress

func init() {
	addr0x1, _ = models.HexToAccountAddress("0x1")

	aptosCoinTypeTag = models.TypeTagStruct{
		Address: addr0x1,
		Module:  "aptos_coin",
		Name:    "AptosCoin",
	}

	escrowAddress = "0x967a8efd835a8c5ac033b467f117c24b5e98216a4791deb8a0e81e7393f35a6e"
	escrowAddr, _ = models.HexToAccountAddress(escrowAddress)
}

func (a *App) WithdrawDeposit(receiverAddress string, deposit *types.Deposit) error {
	aptosClient := client.NewAptosClient("https://fullnode.testnet.aptoslabs.com")
	_, err := client.NewTokenClient(aptosClient)
	if err != nil {
		return fmt.Errorf("failed to create aptos client: %w", err)
	}

	//ctx := context.Background()

	//receiverAccount, err := aptosClient.GetAccount(ctx, receiverAddress)
	//if err != nil {
	//	return fmt.Errorf("failed to get receiver account: %w", err)
	//}

	//escrowAccount, err := aptosClient.GetAccount(ctx, escrowAddress)
	//if err != nil {
	//	return fmt.Errorf("failed to get escrow account: %w", err)
	//}
	//
	//tx := models.Transaction{}
	//err = tx.SetChainID(TestnetChainID).
	//	SetSender(escrowAddress).
	//	SetPayload(getTransferPayload(receiverAccount, deposit.Amount))

	return nil
}

func SendAPT(from, to string, amount uint64) error {
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
