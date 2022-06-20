package smsDeserializer

import (
	"encoding/json"
	"fmt"
	commonDto "github.com/6reduk/smsaeroclient/smsaero/dispatcher/common/model"
	smsDto "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
	"strconv"
)

func UnmarshalSmsDescriptionList(resp *apiResponse.Response) (*smsDto.SmsDescriptionsList, error) {
	var data map[string]json.RawMessage
	err := json.Unmarshal(resp.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal response data, reason: %w", err)
	}

	descriptionList := smsDto.NewSmsDescriptionList()

	for k, v := range data {
		if k == "links" {
			var links *commonDto.Paginator
			err = json.Unmarshal(v, &links)
			if err != nil {
				return nil, fmt.Errorf("unable unmarshal response data %s, reason: %w", v, err)
			}
			descriptionList.SetLinks(links)

			continue
		}

		if k == "totalCount" {
			var countStr string
			err = json.Unmarshal(v, &countStr)
			if err != nil {
				return nil, fmt.Errorf("unable unmarshal response data %s, reason: %w", v, err)
			}

			count, err := strconv.Atoi(countStr)
			if err != nil {
				return nil, fmt.Errorf("can't convert count string to int: %s, reason: %v", countStr, err)
			}
			descriptionList.SetCount(count)

			continue
		}

		var description *smsDto.SmsDescription
		err = json.Unmarshal(v, &description)
		if err != nil {
			return nil, fmt.Errorf("unable unmarshal response data %s, reason: %w", v, err)
		}
		descriptionList.AddDescription(description)
	}

	return descriptionList, nil
}
