package sms

import (
	smsDispatcher "github.com/6reduk/smsaeroclient/internal/dispatcher/sms"
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	testUtil "github.com/6reduk/smsaeroclient/tests/util"
	"github.com/stretchr/testify/suite"
)

type BaseDispatcherShould struct {
	suite.Suite
	config        *testUtil.TestConfig
	client        *httpApiClient.Client
	smsDispatcher *smsDispatcher.SmsDispatcher
}

func (s *BaseDispatcherShould) SetupSuite() {
	testUtil.SkipIntegrationTestIfNeed(s.T())

	s.config = testUtil.GetTestConfigFromEnv()

	s.client = testUtil.GetClientFromTestConfig(s.config)
	s.smsDispatcher = smsDispatcher.NewSmsDispatcher(s.client, s.config.Debug)
}
