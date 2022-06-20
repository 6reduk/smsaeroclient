package smsaero

import (
	authDispatcher "github.com/6reduk/smsaeroclient/smsaero/dispatcher/auth"
	smsDispatcher "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms"
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/request"
)

type ApiFacade struct {
	apiClient httpApiClient.ClientInterface
	sms       smsDispatcher.SmsDispatcherInterface
	auth      authDispatcher.AuthDispatcherInterface
	debug     bool
}

func NewSmsAero(user, password string, debug bool) *ApiFacade {
	config := httpApiClient.GetDefaultConfig()
	cl := httpApiClient.NewClient(config, &apiRequest.BasicAuth{User: user, Password: password})
	facade := &ApiFacade{apiClient: cl, debug: debug}
	facade.init()

	return facade
}

func NewSmsAeroWithConfig(user, password string, debug bool, config *httpApiClient.HttpClientConfig) *ApiFacade {
	cl := httpApiClient.NewClient(config, &apiRequest.BasicAuth{User: user, Password: password})
	facade := &ApiFacade{apiClient: cl, debug: debug}
	facade.init()

	return facade
}

func (s *ApiFacade) Sms() smsDispatcher.SmsDispatcherInterface {
	return s.sms
}

func (s *ApiFacade) SetSmsClient(smsClient smsDispatcher.SmsDispatcherInterface) *ApiFacade {
	s.sms = smsClient
	return s
}

func (s *ApiFacade) Auth() authDispatcher.AuthDispatcherInterface {
	return s.auth
}

func (s *ApiFacade) SetAuthClient(authClient authDispatcher.AuthDispatcherInterface) *ApiFacade {
	s.auth = authClient
	return s
}

func (s *ApiFacade) Debug() bool {
	return s.debug
}

func (s *ApiFacade) SetDebug(debug bool) *ApiFacade {
	s.debug = debug
	return s
}

func (s *ApiFacade) init() {
	s.auth = authDispatcher.NewAuthDispatcher(s.apiClient)
	s.sms = smsDispatcher.NewSmsDispatcher(s.apiClient, s.debug)
}
