package types

type Account struct {
	Nickname  string `json:"nickname"`
	Address   string `json:"address"`
	PubKeyHex string `json:"pub_key_hex"`
	// TODO: add NFT for PFP
}

func NewAccount(nickname, address, pubKeyHex string) *Account {
	return &Account{
		Nickname:  nickname,
		Address:   address,
		PubKeyHex: pubKeyHex,
	}
}
