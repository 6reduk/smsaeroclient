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

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_InvalidSMS_ReturnError() {
	message := requestModelStub.GetInvalidSmsMessage()
	result, err := s.smsDispatcher.Send(context.Background(), message)
	require.Nil(s.T(), result)
	require.NotNil(s.T(), err)
	var apiError *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiError))
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_ValidSMS_ReturnSuccess() {
	message := requestModelStub.GetValidSmsMessage()
	smsList, err := s.smsDispatcher.Send(context.Background(), message)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), smsList)
	assert.Len(s.T(), smsList, 1)
	assert.Greater(s.T(), smsList[0].ID, 0)

	smsStatus, err := s.smsDispatcher.Status(context.Background(), smsList[0].ID)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), smsStatus)
	assert.EqualValues(s.T(), smsList[0], smsStatus)

}
