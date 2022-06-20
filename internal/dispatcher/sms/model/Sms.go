package smsModel

type Sms struct {
	ID           int     `json:"id"`
	From         string  `json:"from"`
	Number       string  `json:"number"`
	Text         string  `json:"text"`
	Status       int     `json:"status"`
	ExtendStatus string  `json:"extendStatus"`
	Channel      string  `json:"channel"`
	Cost         float32 `json:"cost"`
	DateCreate   int     `json:"dateCreate"`
	DateSend     int     `json:"dateSend"`
	DateAnswer   int     `json:"dateAnswer"`
}

type SmsList []*Sms

func (s *Sms) GetStatus() (Status, error) {
	return SmsStatuses.Of(s.Status)
}
