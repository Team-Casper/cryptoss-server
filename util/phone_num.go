package util

import (
	"regexp"
)

const (
	SKT     = "SKT"
	KT      = "KT"
	LGU     = "LGU"
	SKTMVNO = "SKTMVNO"
	KTMVNO  = "KTMVNO"
	LGUMVNO = "LGUMVNO"
)

func IsValidPhoneNumber(num string) bool {
	phoneNumberReg := `^01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$`
	matched, _ := regexp.MatchString(phoneNumberReg, num)
	return matched
}
