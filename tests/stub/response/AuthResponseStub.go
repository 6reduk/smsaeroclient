package responseStub

import (
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
)

func GetAuthSuccessResponse() *apiResponse.Response {
	return &apiResponse.Response{
		Success: true,
		Data:    nil,
		Message: "Successful authorization.",
	}
}

func GetAuthFailResponse() *apiResponse.Response {
	return &apiResponse.Response{
		Success: false,
		Data:    nil,
		Message: "Your request was made with invalid credentials.",
	}
}
