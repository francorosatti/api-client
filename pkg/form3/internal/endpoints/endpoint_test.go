package endpoints

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHttpClient struct {
	mock.Mock
}

func (m *mockHttpClient) Do(r *http.Request) (*http.Response, error) {
	called := m.Called(r)
	return called.Get(0).(*http.Response), called.Error(1)
}

func Test_endpoint_Do(t *testing.T) {
	type fields struct {
		httpClient IHttpClient
		urlFormat  string
		method     string
	}
	tests := []struct {
		name        string
		fields      fields
		opts        []RequestOption
		expectedOut *http.Response
		expectedErr error
	}{
		{
			name:        "given_invalid_parameters_then_return_error",
			fields:      fields{},
			opts:        []RequestOption{WithParam("id", nil)},
			expectedOut: nil,
			expectedErr: errBuildUrl,
		},
		{
			name: "given_invalid_url_then_return_error",
			fields: fields{
				urlFormat: "%%",
			},
			opts:        []RequestOption{WithParam("id", "id")},
			expectedOut: nil,
			expectedErr: errHttpNewRequest,
		},
		{
			name: "given_invalid_client_then_return_error",
			fields: fields{
				httpClient: func() IHttpClient {
					client := &mockHttpClient{}
					client.On("Do", mock.Anything).Return(&http.Response{}, errors.New("mock_error"))
					return client
				}(),
			},
			opts:        []RequestOption{WithParam("id", "id")},
			expectedOut: nil,
			expectedErr: errDoRequest,
		},
		{
			name: "given_valid_client_then_return_error",
			fields: fields{
				httpClient: func() IHttpClient {
					client := &mockHttpClient{}
					client.On("Do", mock.Anything).Return(&http.Response{}, nil)
					return client
				}(),
			},
			opts: []RequestOption{
				WithParam("id", "id"),
				WithBody([]byte("body")),
			},
			expectedOut: &http.Response{},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			e := NewEndpoint(tt.fields.httpClient, tt.fields.urlFormat, tt.fields.method)

			// Act
			got, err := e.Do(tt.opts...)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func Test_endpoint_buildUrl(t *testing.T) {
	tests := []struct {
		name        string
		urlFormat   string
		params      map[string]interface{}
		expectedOut string
		expectedErr error
	}{
		{
			name:        "given_url_without_parameters_then_return_same_url",
			urlFormat:   "https://host:port/path",
			params:      nil,
			expectedOut: "https://host:port/path",
			expectedErr: nil,
		},
		{
			name:        "given_url_with_parameters_then_return_same_url_with_params",
			urlFormat:   "https://host:port/path/{param}",
			params:      map[string]interface{}{"param": "value"},
			expectedOut: "https://host:port/path/value",
			expectedErr: nil,
		},
		{
			name:        "given_url_with_missing_parameters_then_return_same_url",
			urlFormat:   "https://host:port/path/{param}",
			params:      map[string]interface{}{"missing_param": "value"},
			expectedOut: "https://host:port/path/{param}",
			expectedErr: nil,
		},
		{
			name:        "given_url_with_invalid_parameters_then_return_error",
			urlFormat:   "https://host:port/path",
			params:      map[string]interface{}{"param": nil},
			expectedOut: "",
			expectedErr: errSerialiseParamValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			e := endpoint{
				urlFormat: tt.urlFormat,
			}

			// Act
			got, err := e.buildUrl(tt.params)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func Test_serialiseParamValue(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		expectedOut string
		expectedErr error
	}{
		{
			name:        "given_string_parameter_then_serialise_ok",
			value:       "string_param",
			expectedOut: "string_param",
			expectedErr: nil,
		},
		{
			name:        "given_unsupported_parameter_type_then_serialise_ok",
			value:       nil,
			expectedOut: "",
			expectedErr: errUnsupportedParamType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := serialiseParamValue(tt.value)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}
