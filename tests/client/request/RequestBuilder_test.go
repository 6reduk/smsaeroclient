package request

import (
	"fmt"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RequestBuilderShould struct {
	suite.Suite
	builder *apiRequest.RequestBuilder
}

func TestRequestBuilderShould(t *testing.T) {
	suite.Run(t, new(RequestBuilderShould))
}

func (s *RequestBuilderShould) SetupTest() {
	s.builder = apiRequest.NewRequestBuilder()
}

func (s *RequestBuilderShould) TestWhenWrongMethod_returnErr() {
	expectedErrors := []string{
		apiRequest.ErrMethodNotAllowed.Error(),
	}

	req, errors := s.builder.Method(http.MethodConnect).Build()
	assert.Nil(s.T(), req)
	assert.Equal(s.T(), len(expectedErrors), len(errors))
	for i, err := range expectedErrors {
		assert.Equal(s.T(), err, errors[i].Error())
	}
}

func (s *RequestBuilderShould) TestCanSetMethod() {
	allowedMethods := []string{
		http.MethodPost,
		http.MethodGet,
	}

	for _, method := range allowedMethods {
		s.T().Run(fmt.Sprintf("can set method %s", method), func(t *testing.T) {
			localMethod := method // nolint
			req, errors := apiRequest.NewRequestBuilder().Method(localMethod).Build()
			assert.Nil(t, errors)
			assert.Equal(t, localMethod, req.Method)
		})
	}
}

func (s *RequestBuilderShould) TestCanCreateFullPath() {
	url := "http://some-host.somedomain"
	path := "/path/endpoint"
	req, errors := s.builder.URL(url).Path(path).Build()
	assert.Nil(s.T(), errors)
	assert.EqualValues(s.T(), fmt.Sprintf("%s%s", url, path), req.URL.String())
}

func (s *RequestBuilderShould) TestCanSetFormParams() {
	expectedParams := map[string]string{
		"foo":      "bar",
		"intValue": "10",
	}

	s.builder.Method(http.MethodPost)
	for k, v := range expectedParams {
		s.builder.AddFormParam(k, v)
	}

	req, errors := s.builder.Build()
	require.Nil(s.T(), errors)

	for k, v := range expectedParams {
		actualValue := req.FormValue(k)
		assert.Equal(s.T(), v, actualValue)
	}
}

func (s *RequestBuilderShould) TestCanSetHeaders() {
	expectedHeaders := map[string]string{
		"firstHeader":  "valueFirst",
		"secondHeader": "valueSecond",
		"thirdHeader":  "valueThird",
	}

	s.builder.Method(http.MethodPost)
	for key, val := range expectedHeaders {
		s.builder.AddHeader(key, val)
	}

	req, errors := s.builder.Build()
	assert.Nil(s.T(), errors)

	for expectedName, expectedValue := range expectedHeaders {
		value := req.Header.Get(expectedName)
		assert.Equal(s.T(), expectedValue, value)
	}
}

func (s *RequestBuilderShould) TestBasicAuth() {
	user := "username"
	password := "password"

	req, errors := s.builder.Method(http.MethodPost).
		BasicAuth(apiRequest.NewBasicAuth(user, password)).
		Build()
	assert.Nil(s.T(), errors)

	actualUser, actualPassword, ok := req.BasicAuth()
	assert.True(s.T(), ok)
	assert.Equal(s.T(), user, actualUser)
	assert.Equal(s.T(), password, actualPassword)
}

func (s *RequestBuilderShould) TestSetQueryParams() {
	req, errors := s.builder.Method(http.MethodGet).
		AddQueryParam("k1", "v1").
		AddQueryParam("k2", "10").
		AddQueryParam("k3", "v3").
		Path("testPath").
		Build()
	assert.Nil(s.T(), errors)
	assert.Equal(s.T(), "/testPath?k1=v1&k2=10&k3=v3", req.URL.RequestURI())
}
