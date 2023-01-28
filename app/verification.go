package app

import (
	"encoding/json"
	"fmt"
	"github.com/team-casper/cryptoss-server/types"
)

func (a *App) SetVerification(phoneNumber string, verification *types.Verification) error {
	key := types.GetVerificationKey(phoneNumber)

	marshaledVerification, err := json.Marshal(verification)
	if err != nil {
		return fmt.Errorf("failed to marshal verification data: %w", err)
	}

	if err := a.DB.Put(key, marshaledVerification, nil); err != nil {
		return fmt.Errorf("failed to put data: %w", err)
	}

	return nil
}
