package request

import (
	"context"
	"fmt"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type MakeRequestFromShould struct {
	suite.Suite
}

func TestMakeRequestFromShould(t *testing.T) {
	suite.Run(t, &MakeRequestFromShould{})
}

func (s *MakeRequestFromShould) TestMakeRequestFrom_InvalidMethodParameter_ReturnError() {
	parameters := apiRequest.NewRequestParameters().
		SetMethod(http.MethodDelete)

	request, err := apiRequest.MakeRequestFrom(context.Background(), parameters, nil)
	require.Nil(s.T(), request)
	assert.NotNil(s.T(), err)
}

func (s *MakeRequestFromShould) TestMakeRequestFrom_InvalidFormParameterType_ReturnError() {
	invalidParam := uint(2)
	parameters := apiRequest.NewRequestParameters().
		SetMethod(http.MethodDelete).
		AddFormParam("TestKey", invalidParam)

	request, err := apiRequest.MakeRequestFrom(context.Background(), parameters, nil)
	require.Nil(s.T(), request)
	assert.NotNil(s.T(), err)
}

func (s *MakeRequestFromShould) TestMakeRequestFrom_InvalidQueryParameterType_ReturnError() {
	invalidParam := uint(2)
	parameters := apiRequest.NewRequestParameters().
		SetMethod(http.MethodDelete).
		AddQueryParam("TestKey", invalidParam)

	request, err := apiRequest.MakeRequestFrom(context.Background(), parameters, nil)
	require.Nil(s.T(), request)
	assert.NotNil(s.T(), err)
}

func (s *MakeRequestFromShould) TestMakeRequestFrom_ValidParameters_ReturnSuccess() {
	expectedBasicAuth := apiRequest.NewBasicAuth("testUser", "testPassword")
	expectedFormKey := "formKey"
	expectedFormValue := "formValue"
	expectedQueryKey := "formKey"
	expectedQueryValue := "formKey"
	expectedPath := "/endpoint"
	expectedPage := 5
	expectedURI := fmt.Sprintf(
		"/v2%s?%s=%s&%s=%d",
		expectedPath,
		expectedQueryKey,
		expectedQueryValue,
		apiRequest.PaginatorQueryParamName,
		expectedPage,
	)

	parameters := apiRequest.NewRequestParameters().
		SetMethod(http.MethodPost).
		SetPath(expectedPath).
		AddFormParam(expectedFormKey, expectedFormValue).
		AddQueryParam(expectedQueryKey, expectedQueryValue).
		SetPage(expectedPage)

	request, err := apiRequest.MakeRequestFrom(context.Background(), parameters, expectedBasicAuth)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), request)
	require.Equal(s.T(), http.MethodPost, request.Method)
	assert.Equal(s.T(), expectedFormValue, request.FormValue(expectedFormKey))
	user, password, ok := request.BasicAuth()
	require.True(s.T(), ok)
	assert.Equal(s.T(), expectedBasicAuth.User, user)
	assert.Equal(s.T(), expectedBasicAuth.Password, password)

	assert.Equal(s.T(), expectedURI, request.URL.RequestURI())
}
