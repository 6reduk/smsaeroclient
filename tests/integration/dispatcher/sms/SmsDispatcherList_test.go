package sms

import (
	"context"
	"errors"
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	requestModelStub "github.com/6reduk/smsaeroclient/tests/stub/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmsDispatcherListShould struct {
	*BaseDispatcherShould
}

func TestSmsDispatcherListShould(t *testing.T) {
	suite.Run(t, &SmsDispatcherListShould{
		BaseDispatcherShould: &BaseDispatcherShould{},
	})
}

func (s *SmsDispatcherListShould) TestSmsDispatcherSend_InvalidFilter_ReturnError() {
	filter := requestModelStub.GetInvalidFilter()

	result, err := s.smsDispatcher.List(context.Background(), filter)
	require.Nil(s.T(), result)
	require.NotNil(s.T(), err)

	var apiError *httpApiClient.ApiError
	assert.True(s.T(), errors.As(err, &apiError))
}

func (s *SmsDispatcherListShould) TestSmsDispatcherList_ValidFilter_ReturnSuccess() {
	filter := requestModelStub.GetValidFilter()

	descriptionList, err := s.smsDispatcher.List(context.Background(), filter)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), descriptionList)

	assert.Greater(s.T(), len(descriptionList.Descriptions()), 0)
}
