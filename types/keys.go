package types

var (
	VerificationKey = []byte{0x01}
	AccountKey      = []byte{0x02}
)

func GetVerificationKey(phoneNumber string) []byte {
	return append(VerificationKey, []byte(phoneNumber)...)
}

func GetAccountKey(phoneNumber string) []byte {
	return append(AccountKey, []byte(phoneNumber)...)
}
