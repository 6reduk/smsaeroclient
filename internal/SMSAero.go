package smsaero

import (
	smsDispatcher "github.com/6reduk/smsaeroclient/internal/dispatcher/sms"
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
)

type SMSAero struct {
	cl    httpApiClient.ClientInterface
	sms   *smsDispatcher.SmsDispatcher
	debug bool
}

func NewSmsAero(user, password string, debug bool) *SMSAero {
	config := httpApiClient.GetDefaultConfig()
	cl := httpApiClient.NewClient(config, &apiRequest.BasicAuth{User: user, Password: password})
	facade := &SMSAero{cl: cl, debug: debug}
	facade.init()

	return facade
}

func NewSmsAeroWithConfig(user, password string, debug bool, config *httpApiClient.HttpClientConfig) *SMSAero {
	cl := httpApiClient.NewClient(config, &apiRequest.BasicAuth{User: user, Password: password})
	facade := &SMSAero{cl: cl, debug: debug}
	facade.init()

	return facade
}

func (s *SMSAero) SMS() *smsDispatcher.SmsDispatcher {
	return s.sms
}

func (s *SMSAero) init() {
	s.sms = smsDispatcher.NewSmsDispatcher(s.cl, s.debug)
}
