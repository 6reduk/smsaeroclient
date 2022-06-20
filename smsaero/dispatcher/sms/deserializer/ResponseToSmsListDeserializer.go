package smsDeserializer

import (
	"encoding/json"
	"fmt"
	smsModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
)

func UnmarshallSmsList(resp *apiResponse.Response) (smsModel.SmsList, error) {
	var smsList smsModel.SmsList
	err := json.Unmarshal(resp.Data, &smsList)
	if err != nil {
		var sms smsModel.Sms
		err = json.Unmarshal(resp.Data, &sms)
		if err != nil {
			return nil, fmt.Errorf("unable unmarshal sms message results data %s, reason: %w", resp.Data, err)
		}

		smsList = smsModel.SmsList{&sms}
	}

	return smsList, err
}
