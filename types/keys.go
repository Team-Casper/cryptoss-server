package types

var (
	VerificationKey = []byte{0x01}
)

func GetVerificationKey(phoneNumber string) []byte {
	return append(VerificationKey, []byte(phoneNumber)...)
}
