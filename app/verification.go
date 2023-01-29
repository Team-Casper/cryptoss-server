package app

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/types"
)

func (a *App) GetVerification(phoneNumber string) (*types.Verification, error) {
	key := types.GetVerificationKey(phoneNumber)

	has, err := a.DB.Has(key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check if DB has the verification: %w", err)
	}

	if !has {
		return nil, fmt.Errorf("verification not found by the phone number")
	}

	bz, err := a.DB.Get(key, nil)
	if err != nil {
		return nil, fmt.Errorf("error occurs while getting verification: %w", err)
	}

	log.Infof("get verification bz: %v", bz)

	var verification *types.Verification
	if err := json.Unmarshal(bz, &verification); err != nil {
		return nil, fmt.Errorf("failed to unmarshal verification data: %w", err)
	}

	return verification, nil
}

func (a *App) SetVerification(phoneNumber string, verification *types.Verification) error {
	key := types.GetVerificationKey(phoneNumber)

	log.Infof("%v", verification)

	bz, err := json.Marshal(verification)
	if err != nil {
		return fmt.Errorf("failed to marshal verification data: %w", err)
	}

	log.Infof("set verification bz: %v", bz)

	if err := a.DB.Put(key, bz, nil); err != nil {
		return fmt.Errorf("failed to put verification data: %w", err)
	}

	return nil
}
