package smsModel

import commonModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/common/model"

type FilteredSmsList struct {
	smsList SmsList
	links   *commonModel.Paginator
	count   int
}

func NewFilteredList() *FilteredSmsList {
	return &FilteredSmsList{
		smsList: SmsList{},
		links:   &commonModel.Paginator{},
	}
}

func (l *FilteredSmsList) SmsList() SmsList {
	return l.smsList
}

func (l *FilteredSmsList) AddSms(sms *Sms) {
	l.smsList = append(l.smsList, sms)
}

func (l *FilteredSmsList) Links() *commonModel.Paginator {
	return l.links
}

func (l *FilteredSmsList) SetLinks(p *commonModel.Paginator) {
	l.links = p
}

func (l *FilteredSmsList) Count() int {
	return l.count
}

func (l *FilteredSmsList) SetCount(c int) {
	l.count = c
}
