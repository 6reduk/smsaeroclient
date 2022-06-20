package request

import (
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RequestParametersShould struct {
	suite.Suite

	params *apiRequest.RequestParameters
}

func TestRequestParametersShould(t *testing.T) {
	suite.Run(t, &RequestParametersShould{})
}

func (s *RequestParametersShould) SetupTest() {
	s.params = apiRequest.NewRequestParameters()
}

func (s *RequestParametersShould) TestCanSetPath() {
	expectedPath := "balbal"
	s.params.SetPath(expectedPath)
	assert.Equal(s.T(), expectedPath, s.params.Path())
}

func (s *RequestParametersShould) TestCanSetMethod() {
	expectedMethod := http.MethodGet
	s.params.SetMethod(expectedMethod)
	assert.Equal(s.T(), expectedMethod, s.params.Method())
}

func (s *RequestParametersShould) TestCanSetByGet() {
	s.params.ByGet()
	assert.Equal(s.T(), http.MethodGet, s.params.Method())
}

func (s *RequestParametersShould) TestCanSetByPost() {
	s.params.ByPost()
	assert.Equal(s.T(), http.MethodPost, s.params.Method())
}

func (s *RequestParametersShould) TestCanAddFormParam() {
	expectedFormValues := map[string]interface{}{
		"foo": 10,
		"bar": "test",
	}

	for k, v := range expectedFormValues {
		s.params.AddFormParam(k, v)
	}

	actualParams := s.params.FormParams()
	require.Equal(s.T(), len(expectedFormValues), len(actualParams))
	for k, v := range actualParams {
		assert.EqualValues(s.T(), expectedFormValues[k], v)
	}
}

func (s *RequestParametersShould) TestAddQueryParam() {
	expectedQueryValues := map[string]interface{}{
		"foo": 10,
		"bar": "test",
	}

	for k, v := range expectedQueryValues {
		s.params.AddFormParam(k, v)
	}

	actualParams := s.params.FormParams()
	require.Equal(s.T(), len(expectedQueryValues), len(actualParams))
	for k, v := range actualParams {
		assert.EqualValues(s.T(), expectedQueryValues[k], v)
	}
}

func (s *RequestParametersShould) TestCanSetPage() {
	expectedPage := 2
	expectedQueryParamName := apiRequest.PaginatorQueryParamName
	s.params.SetPage(expectedPage)

	actualParams := s.params.QueryParams()
	require.Equal(s.T(), 1, len(actualParams))
	actualPageValue, ok := actualParams[expectedQueryParamName]
	require.True(s.T(), ok)
	assert.EqualValues(s.T(), expectedPage, actualPageValue)
}
