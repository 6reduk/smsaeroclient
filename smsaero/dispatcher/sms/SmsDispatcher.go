package smsDispatcher

import (
	"context"
	"fmt"
	smsDeserializer "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/deserializer"
	smsDto "github.com/6reduk/smsaeroclient/smsaero/dispatcher/sms/model"
	httpClient "github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	apiRequest "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/request"
	apiResponse "github.com/6reduk/smsaeroclient/smsaero/httpApiClient/response"
)

const (
	sendAction   = "send"
	statusAction = "status"
	listAction   = "list"

	sendEndPoint       = "/sms/send"
	testSendEndPoint   = "/sms/testsend"
	statusEndPoint     = "/sms/status"
	testStatusEndPoint = "/sms/teststatus"
	listEndPoint       = "/sms/list"
	testListEndPoint   = "/sms/testlist"
)

var (
	regularPaths = map[string]string{
		sendAction:   sendEndPoint,
		statusAction: statusEndPoint,
		listAction:   listEndPoint,
	}

	debugPaths = map[string]string{
		sendAction:   testSendEndPoint,
		statusAction: testStatusEndPoint,
		listAction:   testListEndPoint,
	}
)

type SmsDispatcher struct {
	cl        httpClient.ClientInterface
	endpoints map[string]string
}

func NewSmsDispatcher(client httpClient.ClientInterface, enableDebug bool) *SmsDispatcher {
	paths := regularPaths
	if enableDebug {
		paths = debugPaths
	}
	return &SmsDispatcher{cl: client, endpoints: paths}
}

func (d *SmsDispatcher) Status(ctx context.Context, id string) (*smsDto.SmsDescription, error) {
	path := d.endpoints[statusAction]
	params := apiRequest.NewRequestParameters().
		SetPath(path).
		ByGet().
		AddQueryParam("id", id)

	resp, err := d.doRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	description, err := smsDeserializer.UnmarshalSmsDescription(resp)
	if err != nil {
		return nil, err
	}

	return description, nil
}

func (d *SmsDispatcher) Send(ctx context.Context, sms *smsDto.SmsMessage) (smsDto.SmsList, error) {
	params := d.prepareParamsForSend(sms)

	resp, err := d.doRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	messageResults, err := smsDeserializer.UnmarshallSmsList(resp)

	if err != nil {
		return nil, err
	}

	return messageResults, nil
}

func (d *SmsDispatcher) List(ctx context.Context, filter *smsDto.SmsFilter) (*smsDto.SmsDescriptionsList, error) {
	params := d.prepareParamsForGetDescriptionList(filter)

	resp, err := d.doRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	messageList, err := smsDeserializer.UnmarshalSmsDescriptionList(resp)

	if err != nil {
		return nil, err
	}

	return messageList, nil
}

func (d *SmsDispatcher) doRequest(ctx context.Context, params *apiRequest.RequestParameters) (*apiResponse.Response, error) {
	resp, err := d.cl.SendRequestFor(ctx, params)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("%w, reason: %s", ErrApiBadRequest, resp.Message)
	}

	if resp.Data == nil {
		return nil, apiRequest.ErrResponseDataFieldIsEmpty
	}

	return resp, nil
}

func (d *SmsDispatcher) prepareParamsForGetDescriptionList(filter *smsDto.SmsFilter) *apiRequest.RequestParameters {
	path := d.endpoints[listAction]
	params := apiRequest.NewRequestParameters().
		SetPath(path).
		ByPost()

	if filter.Text != "" {
		params.AddFormParam("text", filter.Text)
	}

	if filter.Number != "" {
		params.AddFormParam("number", filter.Number)
	}

	if filter.Page > 0 {
		params.AddQueryParam("page", filter.Page)
	}

	return params
}

func (d *SmsDispatcher) prepareParamsForSend(sms *smsDto.SmsMessage) *apiRequest.RequestParameters {
	path := d.endpoints[sendAction]
	params := apiRequest.NewRequestParameters().
		SetPath(path).
		ByPost().
		AddFormParam("sign", sms.Sign()).
		AddFormParam("text", sms.Text())

	if len(sms.Numbers()) > 1 {
		params.AddFormParam("numbers", sms.Numbers())
	} else if len(sms.Numbers()) == 1 {
		params.AddFormParam("number", sms.Numbers()[0])
	}

	if sms.DateSend() > 0 {
		params.AddFormParam("dateSend", sms.DateSend())
	}

	if len(sms.CallbackUrl()) > 0 {
		params.AddFormParam("callbackUrl", sms.CallbackUrl())
	}

	if sms.ShortLink() != 0 {
		params.AddFormParam("shortLink", sms.ShortLink())
	}

	return params
}
