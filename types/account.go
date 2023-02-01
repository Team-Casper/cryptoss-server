package types

import "github.com/portto/aptos-go-sdk/models"

type Account struct {
	Nickname  string         `json:"nickname"`
	Address   string         `json:"address"`
	PubKeyHex string         `json:"pub_key_hex"`
	ProfileID models.TokenID `json:"pfp_id"`
}

func NewAccount(nickname, address, pubKeyHex string) *Account {
	return &Account{
		Nickname:  nickname,
		Address:   address,
		PubKeyHex: pubKeyHex,
	}
}
