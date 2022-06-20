package responseStub

import (
	"encoding/json"
	apiResponse "github.com/6reduk/smsaeroclient/internal/httpApiClient/response"
	responseModelStub "github.com/6reduk/smsaeroclient/tests/stub/model"
	"strconv"
)

func GetSmsFailResponse() *apiResponse.Response {
	return &apiResponse.Response{
		Success: false,
		Data:    nil,
		Message: "Some error occurred",
	}
}

func GetSmsListSuccessResponse() *apiResponse.Response {
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

func GetSmsSuccessResponse() *apiResponse.Response {
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
	messageResult := responseModelStub.GetSmsDescription()

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

func GetSmsDescriptionSuccessResponse() *apiResponse.Response {
	messageListResult := responseModelStub.GetSmsDescriptionList()

	data := map[string]json.RawMessage{}

	var messageResultMarshalled json.RawMessage
	messageResultMarshalled, err := json.Marshal(messageListResult.Descriptions()[0])
	if err != nil {
		panic(err)
	}

	linksMarshalled, err := json.Marshal(messageListResult.Links())
	if err != nil {
		panic(err)
	}
	strconv.Itoa(messageListResult.Count())
	countMarshalled, err := json.Marshal(strconv.Itoa(messageListResult.Count()))
	if err != nil {
		panic(err)
	}

	data["0"] = messageResultMarshalled
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
