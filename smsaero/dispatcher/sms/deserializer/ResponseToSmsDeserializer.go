package smsDeserializer

import (
	"encoding/json"
	"fmt"
	smsModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
)

func UnmarshalStatusResponseToSms(resp *apiResponse.Response) (*smsModel.Sms, error) {
	var description *smsModel.SmsStatus
	err := json.Unmarshal(resp.Data, &description)
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal message result data %s, reason: %w", resp.Data, err)
	}

	sms, err := description.ToSms()
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal response to sms: %w", err)
	}
	return sms, err
}
