package sms

import (
	"context"
	"errors"
	smsDispatcher "github.com/6reduk/smsaeroclient/internal/dispatcher/sms"
	smsDto "github.com/6reduk/smsaeroclient/internal/dispatcher/sms/model"
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
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

type SmsDispatcherListShould struct {
	suite.Suite
	clientMock    *clientMock.HttpClientMock
	smsDispatcher *smsDispatcher.SmsDispatcher
}

func TestSmsDispatcherListShould(t *testing.T) {
	suite.Run(t, &SmsDispatcherListShould{})
}

func (s *SmsDispatcherListShould) SetupTest() {
	s.clientMock = clientMock.GetHttpClientMock()
	s.smsDispatcher = smsDispatcher.NewSmsDispatcher(s.clientMock, false)
}

func (s *SmsDispatcherListShould) TestSmsDispatcherList_BadRequestData_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsFailResponse(), nil)

	messageFilter := smsDto.NewSmsFilterFor("423546", "some text", 3)

	smsMessageListResult, err := s.smsDispatcher.List(context.Background(), messageFilter)
	require.Nil(s.T(), smsMessageListResult)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	assert.True(s.T(), errors.Is(err, smsDispatcher.ErrApiBadRequest))
}

func (s *SmsDispatcherListShould) TestSmsDispatcherList_ApiError_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(nil, responseStub.GetApiErrorCode500())

	messageFilter := smsDto.NewSmsFilterFor("423546", "some text", 3)

	smsMessageListResult, err := s.smsDispatcher.List(context.Background(), messageFilter)
	require.Nil(s.T(), smsMessageListResult)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	var apiErr *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiErr))
}

func (s *SmsDispatcherListShould) TestSmsDispatcherList_ValidRequestData_ReturnSuccess() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetSmsDescriptionSuccessResponse(), nil)

	messageFilter := smsDto.NewSmsFilterFor("423546", "some text", 3)
	smsMessageListResult, err := s.smsDispatcher.List(context.Background(), messageFilter)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), smsMessageListResult)
	expectedList := responseModelStub.GetSmsDescriptionList()
	assert.EqualValues(s.T(), expectedList.Descriptions(), smsMessageListResult.Descriptions())
	assert.EqualValues(s.T(), expectedList.Links(), smsMessageListResult.Links())
	assert.EqualValues(s.T(), expectedList.Count(), smsMessageListResult.Count())
}
