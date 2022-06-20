package sms

import (
	"context"
	"errors"
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	requestModelStub "github.com/6reduk/smsaeroclient/tests/stub/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmsDispatcherSendShould struct {
	*BaseDispatcherShould
}

func TestSmsDispatcherShould(t *testing.T) {
	suite.Run(t, &SmsDispatcherSendShould{
		BaseDispatcherShould: &BaseDispatcherShould{},
	})
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_ValidSMS_ReturnSuccess() {
	sms := requestModelStub.GetValidSmsMessage()
	result, err := s.smsDispatcher.Send(context.Background(), sms)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.Len(s.T(), result, 1)
	assert.Greater(s.T(), result[0].ID, 0)
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_InvalidSMS_ReturnError() {
	sms := requestModelStub.GetInvalidSmsMessage()
	result, err := s.smsDispatcher.Send(context.Background(), sms)
	require.Nil(s.T(), result)
	require.NotNil(s.T(), err)
	var apiError *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiError))
}
