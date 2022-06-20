package responseModelStub

import (
	smsModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
)

func GetSms() *smsModel.Sms {
	return &smsModel.Sms{
		ID:           2,
		From:         "Some sign",
		Number:       "5678",
		Text:         "sms body",
		Status:       3,
		ExtendStatus: "transmitted",
		Channel:      "secret channel",
		Cost:         1.25,
		DateCreate:   456789111,
		DateSend:     3124354645,
		DateAnswer:   12345678,
	}

}

func GetSmsList() smsModel.SmsList {
	messageResult := GetSms()

	return smsModel.SmsList{messageResult}
}

func GetSmsStatus() *smsModel.SmsStatus {
	return &smsModel.SmsStatus{
		ID:           2,
		From:         "Some sign",
		Number:       5678,
		Text:         "sms body",
		Status:       3,
		ExtendStatus: "transmitted",
		Channel:      "secret channel",
		Cost:         "1.25",
		DateCreate:   456789111,
		DateSend:     3124354645,
		DateAnswer:   12345678,
	}
}

func GetSmsDescription() *smsModel.SmsDescription {
	return &smsModel.SmsDescription{
		ID:           2,
		From:         "Some sign",
		Number:       5678,
		Text:         "sms body",
		Status:       3,
		ExtendStatus: "transmitted",
		Channel:      "secret channel",
		Cost:         "1.25",
		DateCreate:   456789111,
		DateSend:     3124354645,
		DateAnswer:   12345678,
	}
}

func GetFilteredList() *smsModel.FilteredSmsList {
	messageList := smsModel.NewFilteredList()
	messageList.AddSms(GetSms())
	messageList.SetLinks(GetPaginator())
	messageList.SetCount(1)

	return messageList
}

func GetFilteredListWithPaginatorOnLastPage() *smsModel.FilteredSmsList {
	messageList := smsModel.NewFilteredList()
	messageList.AddSms(GetSms())
	messageList.SetLinks(GetPaginatorOnLastPage())
	messageList.SetCount(14)

	return messageList
}
