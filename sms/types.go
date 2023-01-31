package sms

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

func CreateSMSMsg(from, to, msgContent string) *SendBody {
	return &SendBody{
		Type:    "SMS",
		From:    from,
		Content: "cryptoss sms service",
		Messages: []*Msg{{
			To:      to,
			Content: msgContent,
		}},
	}
}
