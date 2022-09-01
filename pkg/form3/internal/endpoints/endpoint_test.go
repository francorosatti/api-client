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
	}
	tests := []struct {
		name        string
		fields      fields
		opts        []RequestOption
		expectedOut *http.Response
		expectedErr error
	}{
		{
			name: "given invalid parameters" +
				"when doing request" +
				"then return error",
			opts:        []RequestOption{WithParam("id", nil)},
			expectedOut: nil,
			expectedErr: errBuildUrl,
		},
		{
			name: "given an invalid url" +
				"when doing request" +
				"then return error",
			fields: fields{
				urlFormat: "%%",
			},
			opts:        []RequestOption{WithParam("id", "id")},
			expectedOut: nil,
			expectedErr: errHttpNewRequest,
		},
		{
			name: "given an invalid client" +
				"when doing request" +
				"then return error",
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
			name: "given a valid client" +
				"when doing request" +
				"then return ok",
			fields: fields{
				httpClient: func() IHttpClient {
					client := &mockHttpClient{}
					client.On("Do", mock.Anything).Return(&http.Response{}, nil)
					return client
				}(),
			},
			opts: []RequestOption{
				WithQueryParam("param1", "value1"),
				WithQueryParam("param2", "value2"),
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
			e := NewEndpoint(tt.fields.httpClient, tt.fields.urlFormat, "")

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
			name: "given a url without parameters" +
				"when building url" +
				"then return same url",
			urlFormat:   "https://host:port/path",
			params:      nil,
			expectedOut: "https://host:port/path",
			expectedErr: nil,
		},
		{
			name: "given a url with parameters" +
				"when building url" +
				"then return same url with replaced parameter values",
			urlFormat:   "https://host:port/path/{param}",
			params:      map[string]interface{}{"param": "value"},
			expectedOut: "https://host:port/path/value",
			expectedErr: nil,
		},
		{
			name: "given a url with parameters and non-matching parameters" +
				"when building url" +
				"then return same url without replaced parameter values",
			urlFormat:   "https://host:port/path/{param}",
			params:      map[string]interface{}{"missing_param": "value"},
			expectedOut: "https://host:port/path/{param}",
			expectedErr: nil,
		},
		{
			name: "given a url and invalid parameter values" +
				"when building url" +
				"then return error",
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
			name: "given a string parameter" +
				"when serialising parameter" +
				"then serialise ok",
			value:       "string_param",
			expectedOut: "string_param",
			expectedErr: nil,
		},
		{
			name: "given an integer parameter" +
				"when serialising parameter" +
				"then serialise ok",
			value:       1,
			expectedOut: "1",
			expectedErr: nil,
		},
		{
			name: "given an unsupported parameter type" +
				"when serialising parameter" +
				"then return error",
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
