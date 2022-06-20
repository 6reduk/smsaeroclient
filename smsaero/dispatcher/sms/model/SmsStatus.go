package smsModel

import (
	"fmt"
	"strconv"
)

type SmsStatus struct {
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

func (s *SmsStatus) ToSms() (*Sms, error) {
	cost, err := strconv.ParseFloat(s.Cost, 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert cost to float: %v, reason: %w", s.Cost, err)
	}

	return &Sms{
		ID:           s.ID,
		From:         s.From,
		Number:       strconv.Itoa(s.Number),
		Text:         s.Text,
		Status:       s.Status,
		ExtendStatus: s.ExtendStatus,
		Channel:      s.Channel,
		Cost:         cost,
		DateCreate:   s.DateCreate,
		DateSend:     s.DateSend,
		DateAnswer:   s.DateAnswer,
	}, nil
}
