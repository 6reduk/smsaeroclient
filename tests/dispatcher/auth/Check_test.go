package auth

import (
	"context"
	"errors"
	authDispatcher "github.com/6reduk/smsaeroclient/internal/dispatcher/auth"
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	clientMock "github.com/6reduk/smsaeroclient/tests/mock/client"
	"github.com/6reduk/smsaeroclient/tests/stub"
	responseStub "github.com/6reduk/smsaeroclient/tests/stub/response"
	testUtil "github.com/6reduk/smsaeroclient/tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthDispatcherCheckShould struct {
	suite.Suite
	clientMock     *clientMock.HttpClientMock
	authDispatcher *authDispatcher.AuthDispatcher
}

func TestAuthDispatcherCheckShould(t *testing.T) {
	suite.Run(t, &AuthDispatcherCheckShould{})
}

func (s *AuthDispatcherCheckShould) SetupTest() {
	s.clientMock = clientMock.GetHttpClientMock()
	s.authDispatcher = authDispatcher.NewAuthDispatcher(s.clientMock)
}

func (s *AuthDispatcherCheckShould) TestAuthDispatcherCheck_ClientError_ReturnError() {
	expectedError := stub.GetErrorStub()
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(nil, expectedError)

	isAuthenticated, err := s.authDispatcher.Check(context.Background())
	require.False(s.T(), isAuthenticated)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")

	assert.Equal(s.T(), expectedError, err)
}

func (s *AuthDispatcherCheckShould) TestAuthDispatcherCheck_ApiError_ReturnError() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(nil, responseStub.GetApiErrorCode500())

	isAuthenticated, err := s.authDispatcher.Check(context.Background())
	require.False(s.T(), isAuthenticated)
	require.NotNil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")

	var apiErr *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiErr))
}

func (s *AuthDispatcherCheckShould) TestAuthDispatcherCheck_InvalidCredential_ReturnFail() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetAuthFailResponse(), nil)

	isAuthenticated, err := s.authDispatcher.Check(context.Background())
	require.Nil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	require.False(s.T(), isAuthenticated)
}

func (s *AuthDispatcherCheckShould) TestAuthDispatcherCheck_ValidCredential_ReturnSuccess() {
	s.clientMock.On(
		"SendRequestFor",
		mock.MatchedBy(testUtil.PassAnyObject),
		mock.MatchedBy(testUtil.PassAnyObject),
	).Return(responseStub.GetAuthSuccessResponse(), nil)

	isAuthenticated, err := s.authDispatcher.Check(context.Background())
	require.Nil(s.T(), err)
	s.clientMock.MethodCalled("SendRequestFor")
	require.True(s.T(), isAuthenticated)
}
