package responseStub

import (
	"encoding/json"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
	responseModelStub "github.com/6reduk/smsaeroclient/tests/stub/model"
	"strconv"
)

func GetAnyActionFailResponse() *apiResponse.Response {
	return &apiResponse.Response{
		Success: false,
		Data:    nil,
		Message: "Some error occurred",
	}
}

func GetSmsSendSuccessResponseList() *apiResponse.Response {
	smsList := responseModelStub.GetSmsList()

	marshalled, err := json.Marshal(smsList)
	if err != nil {
		panic(err)
	}

	wrapped := json.RawMessage(marshalled)

	return &apiResponse.Response{
		Success: true,
		Data:    wrapped,
		Message: "All fine",
	}
}

func GetSmsSendSuccessResponse() *apiResponse.Response {
	sms := responseModelStub.GetSms()

	marshalled, err := json.Marshal(sms)
	if err != nil {
		panic(err)
	}

	wrapped := json.RawMessage(marshalled)

	return &apiResponse.Response{
		Success: true,
		Data:    wrapped,
		Message: "All fine",
	}
}

func GetSmsStatusSuccessResponse() *apiResponse.Response {
	messageResult := responseModelStub.GetSmsStatus()

	marshalled, err := json.Marshal(messageResult)
	if err != nil {
		panic(err)
	}

	wrapped := json.RawMessage(marshalled)

	return &apiResponse.Response{
		Success: true,
		Data:    wrapped,
		Message: "All fine.",
	}
}

func GetSmsListSuccessResponse() *apiResponse.Response {
	messageListResult := responseModelStub.GetFilteredList()

	data := map[string]json.RawMessage{}

	var smsMarshalled json.RawMessage
	smsMarshalled, err := json.Marshal(responseModelStub.GetSmsDescription())
	if err != nil {
		panic(err)
	}

	linksMarshalled, err := json.Marshal(responseModelStub.GetPaginator())
	if err != nil {
		panic(err)
	}

	countMarshalled, err := json.Marshal(strconv.Itoa(messageListResult.Count()))
	if err != nil {
		panic(err)
	}

	data["0"] = smsMarshalled
	data["links"] = linksMarshalled
	data["totalCount"] = countMarshalled

	marshalledList, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	wrapped := json.RawMessage(marshalledList)

	return &apiResponse.Response{
		Success: true,
		Data:    wrapped,
		Message: "All fine.",
	}
}

func GetSmsListWithLastPagePaginatorSuccessResponse() *apiResponse.Response {
	messageListResult := responseModelStub.GetFilteredList()

	data := map[string]json.RawMessage{}

	var smsMarshalled json.RawMessage
	smsMarshalled, err := json.Marshal(responseModelStub.GetSmsDescription())
	if err != nil {
		panic(err)
	}

	linksMarshalled, err := json.Marshal(responseModelStub.GetPaginatorOnLastPage())
	if err != nil {
		panic(err)
	}

	countMarshalled, err := json.Marshal(strconv.Itoa(messageListResult.Count()))
	if err != nil {
		panic(err)
	}

	data["0"] = smsMarshalled
	data["links"] = linksMarshalled
	data["totalCount"] = countMarshalled

	marshalledList, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	wrapped := json.RawMessage(marshalledList)

	return &apiResponse.Response{
		Success: true,
		Data:    wrapped,
		Message: "All fine.",
	}
}
