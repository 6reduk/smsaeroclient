package smsDeserializer

import (
	"encoding/json"
	"fmt"
	smsDto "github.com/6reduk/smsaeroclient/internal/dispatcher/sms/model"
	apiResponse "github.com/6reduk/smsaeroclient/internal/httpApiClient/response"
)

func UnmarshallSmsList(resp *apiResponse.Response) (smsDto.SmsList, error) {
	var smsList smsDto.SmsList
	err := json.Unmarshal(resp.Data, &smsList)
	if err != nil {
		var sms smsDto.Sms
		err = json.Unmarshal(resp.Data, &sms)
		if err != nil {
			return nil, fmt.Errorf("unable unmarshal sms message results data %s, reason: %w", resp.Data, err)
		}

		smsList = smsDto.SmsList{&sms}
	}

	return smsList, err
}
