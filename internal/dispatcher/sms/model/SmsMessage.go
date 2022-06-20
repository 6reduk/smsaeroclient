package smsModel

type SmsMessage struct {
	numbers     []string
	sign        string
	text        string
	dateSend    int
	callbackUrl string
	shortLink   int
}

func NewSms() *SmsMessage {
	return &SmsMessage{numbers: make([]string, 0)}
}

func (s *SmsMessage) AddNumbers(numbers []string) {
	s.numbers = append(s.numbers, numbers...)
}

func (s *SmsMessage) Numbers() []string {
	return s.numbers
}

func (s *SmsMessage) SetSign(sign string) {
	s.sign = sign
}

func (s *SmsMessage) Sign() string {
	return s.sign
}

func (s *SmsMessage) SetText(text string) {
	s.text = text
}

func (s *SmsMessage) Text() string {
	return s.text
}

func (s *SmsMessage) SetDateSend(date int) {
	s.dateSend = date
}

func (s *SmsMessage) DateSend() int {
	return s.dateSend
}

func (s *SmsMessage) SetCallbackUrl(url string) {
	s.callbackUrl = url
}

func (s *SmsMessage) CallbackUrl() string {
	return s.callbackUrl
}

func (s *SmsMessage) SetShortLink(flag int) {
	s.shortLink = flag
}

func (s *SmsMessage) EnableShortLink() {
	s.SetShortLink(1)
}

func (s *SmsMessage) DisableShortLink() {
	s.SetShortLink(2)
}

func (s *SmsMessage) ShortLink() int {
	return s.shortLink
}
