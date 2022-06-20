package httpApiClient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	apiRequest "github.com/6reduk/smsaeroclient/internal/httpApiClient/request"
	apiResponse "github.com/6reduk/smsaeroclient/internal/httpApiClient/response"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client struct {
	auth      *apiRequest.BasicAuth
	config    *HttpClientConfig
	transport *http.Transport
	client    *http.Client
}

func NewClient(config *HttpClientConfig, auth *apiRequest.BasicAuth) *Client {
	cl := &Client{
		auth:   auth,
		config: config,
	}

	cl.initTransport()
	cl.initHttpClient()

	return cl
}

func (c *Client) SendRequest(request *http.Request) (*apiResponse.Response, error) {
	httpResp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w, reason: %s", ErrExecutionRequest, err)
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w, reason: %s", ErrReadingBody, err)
	}

	var response *apiResponse.Response
	unmarshallErr := json.Unmarshal(body, &response)
	if unmarshallErr != nil {
		return nil, NewHandleResponseError(err, httpResp.StatusCode, "")
	}

	if !c.isSuccessResponseCode(httpResp.StatusCode) {
		return nil, NewApiError(nil, httpResp.StatusCode, response.Message)
	}

	return response, nil
}

func (c *Client) SendRequestFor(ctx context.Context, parameters *apiRequest.RequestParameters) (*apiResponse.Response, error) {
	request, err := apiRequest.MakeRequestFrom(ctx, parameters, c.auth)
	if err != nil {
		return nil, err
	}

	return c.SendRequest(request)
}

func (c *Client) responseCodeToError(statusCode int) error {
	switch statusCode {

	}

	return ErrExecutionRequest
}

func (c *Client) isSuccessResponseCode(statusCode int) bool {
	return statusCode == 200 || statusCode == 201 || statusCode == 204
}

func (c *Client) initTransport() {
	c.transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(c.config.Timeout) * time.Second,
			KeepAlive: time.Duration(c.config.KeepAlive) * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       time.Duration(c.config.IdleConnTimeout) * time.Second,
		TLSHandshakeTimeout:   time.Duration(c.config.TLSHandshakeTimeout) * time.Second,
		ExpectContinueTimeout: time.Duration(c.config.ExpectContinueTimeout) * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: false},
	}
}

func (c *Client) initHttpClient() {
	c.client = &http.Client{
		Timeout:   time.Second * time.Duration(c.config.Timeout),
		Transport: c.transport,
	}
}
