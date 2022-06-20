package smsModel

type SmsDescription struct {
	ID           int    `json:"id"`
	From         string `json:"from"`
	Number       int    `json:"number"`
	Text         string `json:"text"`
	Status       int    `json:"status"`
	ExtendStatus string `json:"extendStatus"`
	Channel      string `json:"channel"`
	Cost         string `json:"cost"`
	DateCreate   int    `json:"dateCreate"`
	DateSend     int    `json:"dateSend"`
	DateAnswer   int    `json:"dateAnswer"`
}

type SmsDescriptions []*SmsDescription

func (s *SmsDescription) GetStatus() (Status, error) {
	return SmsStatuses.Of(s.Status)
}
