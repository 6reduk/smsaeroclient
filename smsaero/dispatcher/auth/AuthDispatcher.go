package authDispatcher

import (
	"context"
	httpClient "github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/request"
)

const checkAuthPath = "/auth"

type AuthDispatcher struct {
	cl httpClient.ClientInterface
}

func NewAuthDispatcher(client httpClient.ClientInterface) *AuthDispatcher {
	return &AuthDispatcher{cl: client}
}

func (d *AuthDispatcher) Check(ctx context.Context) (bool, error) {
	params := apiRequest.NewRequestParameters().
		SetPath(checkAuthPath).
		ByGet()

	response, err := d.cl.SendRequestFor(ctx, params)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}
