package responseModelStub

import (
	commonModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/common/model"
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

func GetSmsDescriptionList() *smsModel.SmsDescriptionsList {
	messageList := smsModel.NewSmsDescriptionList()
	messageList.AddDescription(GetSmsDescription())
	links := commonModel.NewPaginator()
	links.Self = "/test/self"
	links.Prev = "/test/prev"
	links.Next = "/test/next"
	links.First = "/test/first"
	links.Last = "/test/last"
	messageList.SetLinks(links)
	messageList.SetCount(1)

	return messageList
}
