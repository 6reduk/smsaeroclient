package sms

import (
	"context"
	"errors"
	smsDispatcher "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms"
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	clientMock "github.com/6reduk/smsaeroclient/tests/mock/client"
	responseModelStub "github.com/6reduk/smsaeroclient/tests/stub/model"
	requestModelStub "github.com/6reduk/smsaeroclient/tests/stub/request"
	responseStub "github.com/6reduk/smsaeroclient/tests/stub/response"
	testUtil "github.com/6reduk/smsaeroclient/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmsDispatcherSendShould struct {
	suite.Suite
	clientMock    *clientMock.HttpClientMock
	smsDispatcher *smsDispatcher.SmsDispatcher
}

func TestSmsDispatcherSendShould(t *testing.T) {
	suite.Run(t, &SmsDispatcherSendShould{})
}

func (s *SmsDispatcherSendShould) SetupTest() {
	s.clientMock = clientMock.GetHttpClientMock()
	s.smsDispatcher = smsDispatcher.NewSmsDispatcher(s.clientMock, false)
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_BadRequestData_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsFailResponse(), nil)

	smsMessageResults, err := s.smsDispatcher.Send(context.Background(), requestModelStub.GetValidSmsMessage())
	require.Nil(s.T(), smsMessageResults)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	assert.True(s.T(), errors.Is(err, smsDispatcher.ErrApiBadRequest))
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_ApiError_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(nil, responseStub.GetApiErrorCode500())

	smsMessageResults, err := s.smsDispatcher.Send(context.Background(), requestModelStub.GetValidSmsMessage())
	require.Nil(s.T(), smsMessageResults)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	var apiErr *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiErr))
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_ValidRequestDataForSmsList_ReturnSuccess() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsListSuccessResponse(), nil)

	smsMessageResults, err := s.smsDispatcher.Send(context.Background(), requestModelStub.GetValidSmsMessage())
	require.NoError(s.T(), err)
	require.NotNil(s.T(), smsMessageResults)
	assert.Len(s.T(), smsMessageResults, 1)
	assert.EqualValues(s.T(), responseModelStub.GetSmsList(), smsMessageResults)
}

func (s *SmsDispatcherSendShould) TestSmsDispatcherSend_ValidRequestDataForSingleSms_ReturnSuccess() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsSuccessResponse(), nil)

	smsMessageResults, err := s.smsDispatcher.Send(context.Background(), requestModelStub.GetValidSmsMessage())
	require.NoError(s.T(), err)
	require.NotNil(s.T(), smsMessageResults)
	assert.Len(s.T(), smsMessageResults, 1)
	assert.EqualValues(s.T(), responseModelStub.GetSmsList(), smsMessageResults)
}
