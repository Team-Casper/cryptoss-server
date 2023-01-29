package types

import "time"

type Deposit struct {
	SenderPhoneNumber string    `json:"sender_phone_number"`
	Amount            uint64    `json:"amount"`
	Expiry            time.Time `json:"expiry"`
}

// NewDeposit returns new deposit which is expired in 3 days
func NewDeposit(senderPhoneNumber string, amount uint64) *Deposit {
	return &Deposit{
		SenderPhoneNumber: senderPhoneNumber,
		Amount:            amount,
		Expiry:            time.Now().Add(time.Hour * 72),
	}
}
