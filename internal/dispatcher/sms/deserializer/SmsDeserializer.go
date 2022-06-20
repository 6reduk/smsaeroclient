package smsDeserializer

import (
	"encoding/json"
	"fmt"
	smsDto "github.com/6reduk/smsaeroclient/internal/dispatcher/sms/model"
	apiResponse "github.com/6reduk/smsaeroclient/internal/httpApiClient/response"
)

func UnmarshalSmsDescription(resp *apiResponse.Response) (*smsDto.SmsDescription, error) {
	var description *smsDto.SmsDescription
	err := json.Unmarshal(resp.Data, &description)
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal message result data %s, reason: %w", resp.Data, err)
	}

	return description, err
}
