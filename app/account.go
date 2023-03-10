package app

import (
	"encoding/json"
	"fmt"
	"github.com/team-casper/cryptoss-server/types"
)

func (a *App) GetAccount(phoneNumber string) (*types.Account, error) {
	key := types.GetAccountKey(phoneNumber)

	has, err := a.DB.Has(key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check if DB has the account: %w", err)
	}

	if !has {
		return nil, fmt.Errorf("account not found by the phone number")
	}

	bz, err := a.DB.Get(key, nil)
	if err != nil {
		return nil, fmt.Errorf("error occurs while getting account: %w", err)
	}

	var account *types.Account
	if err := json.Unmarshal(bz, &account); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %w", err)
	}

	return account, nil
}

func (a *App) SetAccount(phoneNumber string, account *types.Account) error {
	key := types.GetAccountKey(phoneNumber)

	bz, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to marshal account data: %w", err)
	}

	if err := a.DB.Put(key, bz, nil); err != nil {
		return fmt.Errorf("failed to put account data: %w", err)
	}

	return nil
}

func (a *App) DeleteAccount(phoneNumber string) error {
	return a.DB.Delete(types.GetAccountKey(phoneNumber), nil)
}
