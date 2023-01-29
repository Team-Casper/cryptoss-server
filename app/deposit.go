package app

import (
	"encoding/json"
	"fmt"
	"github.com/team-casper/cryptoss-server/types"
)

func (a *App) GetDeposit(receiverPhoneNumber string) (*types.Deposit, error) {
	key := types.GetDepositKey(receiverPhoneNumber)

	has, err := a.DB.Has(key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check if DB has the deposit: %w", err)
	}

	if !has {
		return nil, fmt.Errorf("deposit not found by the phone number")
	}

	bz, err := a.DB.Get(key, nil)
	if err != nil {
		return nil, fmt.Errorf("error occurs while getting deposit: %w", err)
	}

	var deposit *types.Deposit
	if err := json.Unmarshal(bz, &deposit); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit data: %w", err)
	}

	return deposit, nil
}

func (a *App) SetDeposit(receiverPhoneNumber string, deposit *types.Deposit) error {
	key := types.GetDepositKey(receiverPhoneNumber)

	bz, err := json.Marshal(deposit)
	if err != nil {
		return fmt.Errorf("failed to marshal deposit data: %w", err)
	}

	if err := a.DB.Put(key, bz, nil); err != nil {
		return fmt.Errorf("failed to put deposit data: %w", err)
	}

	return nil
}

func (a *App) HasClaimableDeposit(phoneNumber string) bool {
	key := types.GetDepositKey(phoneNumber)

	has, _ := a.DB.Has(key, nil)
	return has
}
