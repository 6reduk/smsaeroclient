package clientMock

import (
	"context"
	apiRequest "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/request"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type HttpClientMock struct {
	mock.Mock
}

func GetHttpClientMock() *HttpClientMock {
	return &HttpClientMock{}
}

func (m *HttpClientMock) SendRequest(request *http.Request) (*apiResponse.Response, error) {
	ret := m.Called(request)

	var response *apiResponse.Response
	var err error

	err, ok := ret.Get(1).(error)
	if ok {
		return nil, err
	}

	response, ok = ret.Get(0).(*apiResponse.Response)
	if !ok {
		panic("can't cast return values")
	}

	return response, nil
}

func (m *HttpClientMock) SendRequestFor(ctx context.Context, request *apiRequest.RequestParameters) (*apiResponse.Response, error) {
	ret := m.Called(ctx, request)

	var response *apiResponse.Response
	var err error

	err, ok := ret.Get(1).(error)
	if ok {
		return nil, err
	}

	response, ok = ret.Get(0).(*apiResponse.Response)
	if !ok {
		panic("can't cast return values")
	}

	return response, nil
}
