package types

import (
	"time"
)

type Verification struct {
	Nickname   string    `json:"nickname"`
	TelecoCode string    `json:"teleco_code"`
	Code       string    `json:"code"`
	Expiry     time.Time `json:"expiry"`
}

func (v *Verification) IsExpired() bool {
	return time.Now().After(v.Expiry)
}
