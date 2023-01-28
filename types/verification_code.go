package types

import (
	"time"
)

type Verification struct {
	Nickname   string
	TelecoCode string
	Code       string
	Expiry     time.Time
}
