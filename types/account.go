package types

type Account struct {
	Nickname string `json:"nickname"`
	Address  string `json:"address"`
	// TODO: add NFT for PFP
}

func NewAccount(nickname, address string) *Account {
	return &Account{
		Nickname: nickname,
		Address:  address,
	}
}
