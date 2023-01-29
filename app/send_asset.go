package app

import (
	"github.com/portto/aptos-go-sdk/client"
	"github.com/team-casper/cryptoss-server/types"
)

func (a *App) WithdrawDeposit(deposit *types.Deposit) error {
	aptosClient := client.NewAptosClient("https://fullnode.testnet.aptoslabs.com")
	_, err := client.NewTokenClient(aptosClient)
	if err != nil {
		panic(err)
	}

	return nil
}

func SendAPT(from, to string, amount uint64) error {
	return nil
}
