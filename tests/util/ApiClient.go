package testUtil

import (
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/request"
)

func GetClientFromTestConfig(config *TestConfig) *httpApiClient.Client {
	basicAuth := apiRequest.NewBasicAuth(config.User, config.Password)
	cl := httpApiClient.NewClient(httpApiClient.GetDefaultConfig(), basicAuth)
	return cl
}
