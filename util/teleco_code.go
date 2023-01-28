package util

import "strings"

func IsValidTelecoCode(code string) bool {
	switch strings.ToUpper(code) {
	case SKT, KT, LGU, SKTMVNO, KTMVNO, LGUMVNO:
		return true
	default:
		return false
	}

}
