package sms

import "fmt"

type Msg struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type SendBody struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	Content  string `json:"content"`
	Messages []*Msg `json:"messages"`
}

func CreateSMSMsg(from, to, code string) *SendBody {
	return &SendBody{
		Type:    "SMS",
		From:    from,
		Content: "cryptoss verification code",
		Messages: []*Msg{{
			To:      to,
			Content: fmt.Sprintf("[Cryptoss] 인증번호 [%s]를 5분 안에 입력해주세요. ", code),
		}},
	}
}
