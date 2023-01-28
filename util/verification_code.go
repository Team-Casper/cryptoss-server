package util

import "crypto/rand"

const codeChars = "1234567890"

func GenVerificationCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(codeChars)
	for i := 0; i < length; i++ {
		buffer[i] = codeChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
