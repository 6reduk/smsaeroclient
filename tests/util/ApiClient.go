package testUtil

import (
	"github.com/6reduk/smsaeroclient/internal/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
)

func GetClientFromTestConfig(config *TestConfig) *httpApiClient.Client {
	basicAuth := apiRequest.NewBasicAuth(config.User, config.Password)
	cl := httpApiClient.NewClient(httpApiClient.GetDefaultConfig(), basicAuth)
	return cl
}
