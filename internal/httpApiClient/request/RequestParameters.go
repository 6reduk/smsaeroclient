package apiRequest

import "net/http"

const PaginatorQueryParamName = "page"

type RequestParameters struct {
	path        string
	method      string
	formParams  map[string]interface{}
	queryParams map[string]interface{}
}

func NewRequestParameters() *RequestParameters {
	return &RequestParameters{
		formParams:  make(map[string]interface{}),
		queryParams: make(map[string]interface{}),
	}
}

func (r *RequestParameters) Path() string {
	return r.path
}

func (r *RequestParameters) SetPath(path string) *RequestParameters {
	r.path = path
	return r
}

func (r *RequestParameters) SetMethod(method string) *RequestParameters {
	r.method = method
	return r
}

func (r *RequestParameters) ByGet() *RequestParameters {
	r.SetMethod(http.MethodGet)
	return r
}

func (r *RequestParameters) ByPost() *RequestParameters {
	r.SetMethod(http.MethodPost)
	return r
}

func (r *RequestParameters) Method() string {
	return r.method
}

func (r *RequestParameters) AddFormParam(k string, v interface{}) *RequestParameters {
	r.formParams[k] = v
	return r
}

func (r *RequestParameters) FormParams() map[string]interface{} {
	return r.formParams
}

func (r *RequestParameters) AddQueryParam(k string, v interface{}) *RequestParameters {
	r.queryParams[k] = v
	return r
}

func (r *RequestParameters) QueryParams() map[string]interface{} {
	return r.queryParams
}

func (r *RequestParameters) SetPage(pageNumber int) *RequestParameters {
	r.AddQueryParam(PaginatorQueryParamName, pageNumber)
	return r
}
