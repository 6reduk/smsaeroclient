package apiRequest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilder struct {
	method           string
	supportedMethods map[string]bool
	url              string
	path             string
	formParams       map[string][]string
	queryParams      map[string][]string
	headers          map[string]string
	basicAuth        *BasicAuth
	ctx              context.Context
	errors           []error
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		headers:     make(map[string]string),
		formParams:  make(map[string][]string),
		queryParams: make(map[string][]string),
		supportedMethods: map[string]bool{
			http.MethodPost: true,
			http.MethodGet:  true,
		},
	}
}

func (r *RequestBuilder) Build() (*http.Request, []error) {
	if r.hasErrors() {
		return nil, r.errors
	}

	if r.ctx == nil {
		r.ctx = context.Background()
	}
	body := r.getBody()

	request, err := http.NewRequestWithContext(r.ctx, r.method, r.createFullURL(), body)
	if err != nil {
		return nil, []error{err}
	}

	r.fillDefaultHeaders(request)
	r.fillHeaders(request)
	r.fillBasicAuth(request)
	r.fillFormParams(request)
	r.fillQueryParams(request)

	return request, nil
}

func (r *RequestBuilder) Method(method string) *RequestBuilder {
	if _, ok := r.supportedMethods[method]; !ok {
		r.addError(ErrMethodNotAllowed)
		return r
	}
	r.method = method
	return r
}

func (r *RequestBuilder) URL(url string) *RequestBuilder {
	r.url = strings.TrimRight(url, "/")
	return r
}

func (r *RequestBuilder) Path(path string) *RequestBuilder {
	r.path = strings.TrimLeft(path, "/")
	return r
}

func (r *RequestBuilder) Context(ctx context.Context) *RequestBuilder {
	r.ctx = ctx
	return r
}

func (r *RequestBuilder) AddFormParam(key string, value string) *RequestBuilder {
	//r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	r.formParams[key] = append(r.formParams[key], value)
	return r
}

func (r *RequestBuilder) AddQueryParam(key, value string) *RequestBuilder {
	r.queryParams[key] = append(r.queryParams[key], value)
	return r
}

func (r *RequestBuilder) AddHeader(key, value string) *RequestBuilder {
	r.headers[key] = value
	return r
}

func (r *RequestBuilder) BasicAuth(auth *BasicAuth) *RequestBuilder {
	r.basicAuth = auth
	return r
}

func (r *RequestBuilder) addError(err error) {
	r.errors = append(r.errors, err)
}

func (r *RequestBuilder) hasErrors() bool {
	return len(r.errors) > 0
}

func (r *RequestBuilder) createFullURL() string {
	if r.path == "" {
		return r.url
	}

	return fmt.Sprintf("%s/%s", r.url, r.path)
}

func (r *RequestBuilder) fillBasicAuth(request *http.Request) {
	if r.basicAuth == nil {
		return
	}

	request.SetBasicAuth(r.basicAuth.User, r.basicAuth.Password)
}

func (r *RequestBuilder) fillDefaultHeaders(request *http.Request) {
	request.Header.Set("Accept", "application/json")
}

func (r *RequestBuilder) fillHeaders(request *http.Request) {
	for key, val := range r.headers {
		request.Header.Set(key, val)
	}
}

func (r *RequestBuilder) getBody() io.Reader {
	if len(r.formParams) == 0 {
		return bytes.NewBuffer([]byte{})
	}

	urlValues := url.Values{}
	for key, values := range r.formParams {
		for _, val := range values {
			urlValues.Add(key, val)
		}
	}

	return strings.NewReader(urlValues.Encode())
}

func (r *RequestBuilder) fillFormParams(request *http.Request) {

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

func (r *RequestBuilder) fillQueryParams(request *http.Request) {
	if len(r.queryParams) == 0 {
		return
	}
	query := request.URL.Query()
	for key, values := range r.queryParams {
		for _, val := range values {
			query.Add(key, val)
		}
	}
	request.URL.RawQuery = query.Encode()
	println(request)
}
