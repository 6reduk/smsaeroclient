package smsModel

import "fmt"

var (
	ErrInvalidPhone           = fmt.Errorf("invalid phone value")
	ErrPhoneNumberIsRequired  = fmt.Errorf("phone number is required")
	ErrSignIsRequired         = fmt.Errorf("sign is required")
	ErrTextIsRequired         = fmt.Errorf("text is required")
	ErrUnknownMessageStatusID = fmt.Errorf("unknown sms status id")
)
