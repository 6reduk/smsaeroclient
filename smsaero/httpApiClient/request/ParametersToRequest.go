package apiRequest

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const smsAeroURL = "https://gate.smsaero.ru/v2"

//const smsAeroURL = "http://localhost:8080"

func MakeRequestFrom(ctx context.Context, requestParams *RequestParameters, basicAuth *BasicAuth) (*http.Request, error) {
	builder := NewRequestBuilder().
		URL(smsAeroURL).
		Context(ctx)

	if requestParams.Path() == "" {
		return nil, ErrPathIsRequired
	}
	builder.Path(requestParams.Path())

	if requestParams.Method() == "" {
		return nil, ErrMethodIsRequired
	}
	builder.Method(requestParams.Method())

	err := fillBuilderFormFrom(builder, requestParams.FormParams())
	if err != nil {
		return nil, err
	}

	err = fillBuilderQueryFrom(builder, requestParams.QueryParams())
	if err != nil {
		return nil, err
	}

	if basicAuth != nil {
		builder.BasicAuth(basicAuth)
	}

	request, errors := builder.Build()

	if errors != nil {
		wrappedErrors := wrapErrors(errors)
		return nil, wrappedErrors
	}

	return request, nil
}

func fillBuilderFormFrom(builder *RequestBuilder, params map[string]interface{}) error {
	converted, err := convertValuesToString(params)
	if err != nil {
		return fmt.Errorf("%w, reason: %s", ErrUnableTransformFormParam, err)
	}

	for paramName, paramList := range converted {
		for _, v := range paramList {
			builder.AddFormParam(paramName, v)
		}
	}

	return nil
}

func fillBuilderQueryFrom(builder *RequestBuilder, params map[string]interface{}) error {
	converted, err := convertValuesToString(params)
	if err != nil {
		return fmt.Errorf("%w, reason: %s", ErrUnableTransformQueryParam, err)
	}

	for paramName, paramList := range converted {
		for _, v := range paramList {
			builder.AddQueryParam(paramName, v)
		}
	}

	return nil
}

func convertValuesToString(params map[string]interface{}) (map[string][]string, error) {
	converted := make(map[string][]string)
	for paramName, paramValue := range params {
		switch value := paramValue.(type) {
		case int:
			v := strconv.Itoa(value)
			converted[paramName] = append(converted[paramName], v)
		case int64:
			v := strconv.FormatInt(value, 10)
			converted[paramName] = append(converted[paramName], v)
		case []int:
			for _, vInt := range value {
				v := strconv.Itoa(vInt)
				converted[paramName] = append(converted[paramName], v)
			}
		case float32:
			v := fmt.Sprintf("%f", value)
			converted[paramName] = append(converted[paramName], v)
		case []float32:
			for _, vFloat := range value {
				v := fmt.Sprintf("%f", vFloat)
				converted[paramName] = append(converted[paramName], v)
			}
		case float64:
			v := fmt.Sprintf("%f", value)
			converted[paramName] = append(converted[paramName], v)
		case []float64:
			for _, vFloat := range value {
				v := fmt.Sprintf("%f", vFloat)
				converted[paramName] = append(converted[paramName], v)
			}
		case string:
			converted[paramName] = append(converted[paramName], value)
		case []string:
			for _, v := range value {
				converted[paramName] = append(converted[paramName], v)
			}
		case bool:
			v := fmt.Sprintf("%v", value)
			converted[paramName] = append(converted[paramName], v)
		default:
			return nil, fmt.Errorf("unsupported param type %s", paramValue)
		}
	}

	return converted, nil
}

func wrapErrors(errors []error) error {
	errReason := ""
	for _, err := range errors {
		errReason += fmt.Sprintf(" <%s>", err.Error())
	}

	return fmt.Errorf("unable make request from parameters: %s", errReason)
}
