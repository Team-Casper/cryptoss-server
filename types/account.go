package types

import "github.com/portto/aptos-go-sdk/models"

type Account struct {
	Nickname  string         `json:"nickname"`
	Address   string         `json:"address"`
	ProfileID models.TokenID `json:"pfp_id"`
}

func NewAccount(nickname, address string) *Account {
	return &Account{
		Nickname: nickname,
		Address:  address,
	}
}
