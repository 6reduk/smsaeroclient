package smsDispatcher

import (
	"context"
	smsDto "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
)

type SmsDispatcherInterface interface {
	Status(ctx context.Context, id int) (*smsDto.Sms, error)
	Send(ctx context.Context, sms *smsDto.SmsMessage) (smsDto.SmsList, error)
	List(ctx context.Context, filter *smsDto.SmsFilter) (*smsDto.FilteredSmsList, error)
}
