package sms

import (
	"fmt"
	"math"
)

func GetVerificationMsgContent(code string) string {
	return fmt.Sprintf("[Cryptoss] 인증번호 [%s]를 5분 안에 입력해주세요. ", code)
}

func GetInvitationMsgContent(senderNickname, senderPhoneNumber string, amount uint64) string {
	aptAmount := float64(amount) / math.Pow10(8)
	return fmt.Sprintf("[Cryptoss] %s(%s) 님이 %.4f APT를 보냈어요. 지금 확인해보세요!\n\ncryptoss.xyz", senderNickname, senderPhoneNumber, aptAmount)
}
