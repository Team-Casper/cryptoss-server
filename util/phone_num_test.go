package util_test

import (
	"github.com/stretchr/testify/require"
	"github.com/team-casper/cryptoss-server/util"
	"testing"
)

func TestValidPhoneNumber(t *testing.T) {
	num := "01012345678"
	isValid := util.IsValidPhoneNumber(num)
	require.True(t, isValid)
}

func TestInvalidPhoneNumber(t *testing.T) {
	num := "01289897979"
	isValid := util.IsValidPhoneNumber(num)
	require.False(t, isValid)
}
