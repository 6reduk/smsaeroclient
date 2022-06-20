package sms

import (
	"context"
	"errors"
	smsDispatcher "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms"
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	clientMock "github.com/6reduk/smsaeroclient/tests/mock/client"
	responseModelStub "github.com/6reduk/smsaeroclient/tests/stub/model"
	responseStub "github.com/6reduk/smsaeroclient/tests/stub/response"
	testUtil "github.com/6reduk/smsaeroclient/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmsDispatcherStatusShould struct {
	suite.Suite
	clientMock    *clientMock.HttpClientMock
	smsDispatcher *smsDispatcher.SmsDispatcher
}

func TestSmsDispatcherStatusShould(t *testing.T) {
	suite.Run(t, &SmsDispatcherStatusShould{})
}

func (s *SmsDispatcherStatusShould) SetupTest() {
	s.clientMock = clientMock.GetHttpClientMock()
	s.smsDispatcher = smsDispatcher.NewSmsDispatcher(s.clientMock, false)
}

func (s *SmsDispatcherStatusShould) TestSmsDispatcherStatus_BadRequestData_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsFailResponse(), nil)

	smsMessageResults, err := s.smsDispatcher.Status(context.Background(), "1")
	require.Nil(s.T(), smsMessageResults)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	assert.True(s.T(), errors.Is(err, smsDispatcher.ErrApiBadRequest))
}

func (s *SmsDispatcherStatusShould) TestSmsDispatcherStatus_ApiError_ReturnError() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(nil, responseStub.GetApiErrorCode500())

	smsMessageResults, err := s.smsDispatcher.Status(context.Background(), "1")
	require.Nil(s.T(), smsMessageResults)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	var apiErr *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiErr))
}

func (s *SmsDispatcherStatusShould) TestAuthDispatcher_ValidRequestData_ReturnSuccess() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsStatusSuccessResponse(), nil)

	smsMessageResult, err := s.smsDispatcher.Status(context.Background(), "1")
	require.NoError(s.T(), err)
	require.NotNil(s.T(), smsMessageResult)
	assert.EqualValues(s.T(), responseModelStub.GetSmsDescription(), smsMessageResult)
}
