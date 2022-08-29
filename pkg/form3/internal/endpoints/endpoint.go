package endpoints

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	RequestOption func(*requestOptions)

	IEndpoint interface {
		Do(opts ...RequestOption) (*http.Response, error)
	}

	IHttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	endpoint struct {
		httpClient IHttpClient
		urlFormat  string
		method     string
	}

	requestOptions struct {
		params map[string]interface{}
		body   []byte
	}
)

func NewEndpoint(client IHttpClient, url string, method string) IEndpoint {
	return endpoint{
		httpClient: client,
		urlFormat:  url,
		method:     method,
	}
}

func WithParam(key string, value interface{}) RequestOption {
	return func(options *requestOptions) {
		options.params[key] = value
	}
}

func WithBody(body []byte) RequestOption {
	return func(options *requestOptions) {
		options.body = body
	}
}

func (e endpoint) Do(opts ...RequestOption) (*http.Response, error) {
	options := defaultRequestOptions()

	for _, opt := range opts {
		opt(&options)
	}

	url, err := e.buildUrl(options.params)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errBuildUrl, err)
	}

	var bodyReader io.Reader
	if options.body != nil {
		bodyReader = bytes.NewBuffer(options.body)
	}

	req, err := http.NewRequest(e.method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errHttpNewRequest, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errDoRequest, err)
	}

	return res, nil
}

func defaultRequestOptions() requestOptions {
	return requestOptions{
		params: make(map[string]interface{}),
		body:   nil,
	}
}

func (e endpoint) buildUrl(params map[string]interface{}) (string, error) {
	url := e.urlFormat

	for key, value := range params {
		strValue, err := serialiseParamValue(value)
		if err != nil {
			return "", errSerialiseParamValue
		}

		url = strings.ReplaceAll(url, fmt.Sprintf("{%s}", key), strValue)
	}

	return url, nil
}

func serialiseParamValue(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		return value.(string), nil
	default:
		return "", errUnsupportedParamType
	}
}
