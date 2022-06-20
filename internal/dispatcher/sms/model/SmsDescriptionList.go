package smsModel

import commonModel "github.com/6reduk/smsaeroclient/internal/dispatcher/common/model"

type SmsDescriptionsList struct {
	smsDescriptions SmsDescriptions
	links           *commonModel.Paginator
	count           int
}

func NewSmsDescriptionList() *SmsDescriptionsList {
	return &SmsDescriptionsList{
		smsDescriptions: SmsDescriptions{},
		links:           &commonModel.Paginator{},
	}
}

func (l *SmsDescriptionsList) Descriptions() SmsDescriptions {
	return l.smsDescriptions
}

func (l *SmsDescriptionsList) AddDescription(message *SmsDescription) {
	l.smsDescriptions = append(l.smsDescriptions, message)
}

func (l *SmsDescriptionsList) Links() *commonModel.Paginator {
	return l.links
}

func (l *SmsDescriptionsList) SetLinks(p *commonModel.Paginator) {
	l.links = p
}

func (l *SmsDescriptionsList) Count() int {
	return l.count
}

func (l *SmsDescriptionsList) SetCount(c int) {
	l.count = c
}
