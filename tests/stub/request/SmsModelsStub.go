package requestModelStub

import (
	"fmt"
	smsDto "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
)

const (
	validPhoneNumber        = "+7-999-999-99-99"
	validCleanedPhoneNumber = "79999999999"
	invalidPhoneNumber      = "sdfghjkloiuytr"
)

func GetValidSmsMessage() *smsDto.SmsMessage {
	sms, errors := smsDto.NewSmsMessageBuilder().
		Number(validPhoneNumber).
		Sign("Sms aero").
		Text("test text").
		SendAtDate(1781882934).
		CallbackUrl("http://call.back/test").
		Build()

	if len(errors) > 0 {
		panic(fmt.Sprintf("error occured during build sms: %v", errors))
	}

	return sms
}

func GetInvalidSmsMessage() *smsDto.SmsMessage {
	sms := smsDto.NewSmsMessage()
	sms.AddNumbers([]string{invalidPhoneNumber})
	sms.SetSign("foo")
	sms.SetText("bar")

	return sms
}

func GetInvalidFilter() *smsDto.SmsFilter {
	filter := smsDto.NewSmsFilter()
	filter.Number = "invalidPhoneNumber"
	return filter
}

func GetValidFilter() *smsDto.SmsFilter {
	filter := smsDto.NewSmsFilter()
	filter.Number = validCleanedPhoneNumber
	return filter
}
