package smsModel

type SmsFilter struct {
	Number string
	Text   string
	Page   int
}

func NewSmsFilter() *SmsFilter {
	return &SmsFilter{}
}

func NewSmsFilterFor(number, text string, page int) *SmsFilter {
	return &SmsFilter{
		Number: number,
		Text:   text,
		Page:   page,
	}
}
