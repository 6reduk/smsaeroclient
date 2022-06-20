package httpApiClient

import (
	"context"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
	apiResponse "github.com/6reduk/smsaeroclient/internal/httpApiClient/response"
	"net/http"
)

type ClientInterface interface {
	SendRequest(request *http.Request) (*apiResponse.Response, error)
	SendRequestFor(ctx context.Context, parameters *apiRequest.RequestParameters) (*apiResponse.Response, error)
}
