package smsModel

import (
	"fmt"
	"regexp"
	"strings"
)

const checkPhoneRegex = `^\d+$`

type SmsMessageBuilder struct {
	numbers     []string
	sign        string
	text        string
	dateSend    int
	callbackUrl string
	shortLink   bool
	errors      []error
}

func NewSmsMessageBuilder() *SmsMessageBuilder {
	return &SmsMessageBuilder{
		numbers: make([]string, 0),
		errors:  make([]error, 0),
	}
}

func (b *SmsMessageBuilder) Build() (*SmsMessage, []error) {
	b.checkRequiredFields()

	if b.hasErrors() {
		return nil, b.errors
	}

	sms := NewSmsMessage()
	sms.AddNumbers(b.numbers)
	sms.SetSign(b.sign)
	sms.SetText(b.text)
	sms.SetDateSend(b.dateSend)
	sms.SetCallbackUrl(b.callbackUrl)
	if b.shortLink {
		sms.EnableShortLink()
	}

	return sms, nil
}

func (b *SmsMessageBuilder) Number(number string) *SmsMessageBuilder {
	cleanedNumber, err := b.cleanNumber(number)
	if err != nil {
		b.addError(err)
	} else {
		b.numbers = append(b.numbers, cleanedNumber)
	}
	return b
}

func (b *SmsMessageBuilder) Numbers(numbers []string) *SmsMessageBuilder {
	for _, number := range numbers {
		cleanedNumber, err := b.cleanNumber(number)
		if err != nil {
			b.addError(err)
		}

		b.numbers = append(b.numbers, cleanedNumber)
	}

	return b
}

func (b *SmsMessageBuilder) Sign(sign string) *SmsMessageBuilder {
	b.sign = sign
	return b
}

func (b *SmsMessageBuilder) Text(text string) *SmsMessageBuilder {
	b.text = text
	return b
}

func (b *SmsMessageBuilder) SendAtDate(date int) *SmsMessageBuilder {
	b.dateSend = date
	return b
}

func (b *SmsMessageBuilder) CallbackUrl(url string) *SmsMessageBuilder {
	b.callbackUrl = url
	return b
}

func (b *SmsMessageBuilder) UseShortLink(flag bool) *SmsMessageBuilder {
	b.shortLink = flag
	return b
}

func (b *SmsMessageBuilder) checkRequiredFields() {
	if len(b.numbers) == 0 {
		b.addError(ErrPhoneNumberIsRequired)
	}

	if len(b.sign) == 0 {
		b.addError(ErrSignIsRequired)
	}

	if len(b.text) == 0 {
		b.addError(ErrTextIsRequired)
	}
}

func (b *SmsMessageBuilder) cleanNumber(number string) (string, error) {

	cleanedNumber := strings.TrimLeft(number, "+")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "-", "")
	matched, err := regexp.Match(checkPhoneRegex, []byte(cleanedNumber))
	if err != nil {
		return number, fmt.Errorf("can't check phone: %s, %w", number, err)
	}

	if !matched {
		return number, fmt.Errorf("invalid phone value: %s, %w", number, ErrInvalidPhone)
	}

	return cleanedNumber, nil
}

func (b *SmsMessageBuilder) addError(err error) {
	b.errors = append(b.errors, err)
}

func (b *SmsMessageBuilder) hasErrors() bool {
	return len(b.errors) > 0
}
