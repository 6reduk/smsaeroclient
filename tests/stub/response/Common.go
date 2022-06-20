package responseStub

import (
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	"net/http"
)

func GetApiErrorCode500() *httpApiClient.ApiError {
	return httpApiClient.NewApiError(nil, http.StatusInternalServerError, "Not deleted")
}
