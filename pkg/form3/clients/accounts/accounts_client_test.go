package accounts

import (
	"errors"
	"github.com/francorosatti/form3-api-client/pkg/form3/internal/endpoints"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
)

type endpointMock struct {
	mock.Mock
}

func (m *endpointMock) Do(opts ...endpoints.RequestOption) (*http.Response, error) {
	called := m.Called(opts)
	return called.Get(0).(*http.Response), called.Error(1)
}

func TestNewAccountClient(t *testing.T) {
	// Arrange
	baseUrl := "baseUrl"

	// Act
	clientIntf := NewAccountClient(baseUrl)

	// Assert
	client := clientIntf.(accountClient)
	_, exists := client.endpoints[_endpointCreateAccount]
	assert.True(t, exists)
	_, exists = client.endpoints[_endpointFetchAccount]
	assert.True(t, exists)
	_, exists = client.endpoints[_endpointDeleteAccount]
	assert.True(t, exists)
}

func Test_accountClient_CreateAccount(t *testing.T) {
	type fields struct {
		endpoint endpoints.IEndpoint
	}
	tests := []struct {
		name        string
		fields      fields
		expectedOut models.Account
		expectedErr error
	}{
		{
			name: "given_endpoint_status_bad_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 400,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: ErrAccountBadRequest,
		},
		{
			name: "given_endpoint_invalid_json_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader("}")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: errResponseUnmarshal,
		},
		{
			name: "given_endpoint_status_ok_then_return_model",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader(`{"data":{"attributes":{"bank_id":"bank_id"},"id":"id"}}`)),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: nil,
			expectedOut: models.Account{
				Data: &models.AccountData{
					ID: "id",
					Attributes: &models.AccountAttributes{
						BankID: "bank_id",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointCreateAccount: tt.fields.endpoint,
				},
			}
			got, err := client.CreateAccount(models.Account{})
			require.True(t, errors.Is(err, tt.expectedErr))
			require.Equal(t, tt.expectedOut, got)
		})
	}
}

func Test_accountClient_FetchAccount(t *testing.T) {
	type fields struct {
		endpoint endpoints.IEndpoint
	}
	tests := []struct {
		name        string
		fields      fields
		expectedOut models.Account
		expectedErr error
	}{
		{
			name: "given_endpoint_status_not_found_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 404,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: ErrAccountNotFound,
		},
		{
			name: "given_endpoint_status_ok_then_return_model",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader(`{"data":{"attributes":{"bank_id":"bank_id"},"id":"id"}}`)),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: nil,
			expectedOut: models.Account{
				Data: &models.AccountData{
					ID: "id",
					Attributes: &models.AccountAttributes{
						BankID: "bank_id",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointFetchAccount: tt.fields.endpoint,
				},
			}
			got, err := client.FetchAccount("any_id")
			require.True(t, errors.Is(err, tt.expectedErr))
			require.Equal(t, tt.expectedOut, got)
		})
	}
}

func Test_accountClient_requestCreateAccount(t *testing.T) {
	type fields struct {
		endpoint endpoints.IEndpoint
	}
	tests := []struct {
		name        string
		fields      fields
		expectedOut []byte
		expectedErr error
	}{
		{
			name: "given_endpoint_fail_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{}, errors.New("mock_error"))
					return endpoint
				}(),
			},
			expectedErr: errDoRequest,
		},
		{
			name: "given_endpoint_status_bad_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 400,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: ErrAccountBadRequest,
		},
		{
			name: "given_endpoint_status_ok_then_return_ok",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader("{}")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: nil,
			expectedOut: []byte("{}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointCreateAccount: tt.fields.endpoint,
				},
			}
			got, err := client.requestCreateAccount([]byte("anything"))
			require.True(t, errors.Is(err, tt.expectedErr))
			require.Equal(t, tt.expectedOut, got)
		})
	}
}

func Test_accountClient_requestDeleteAccount(t *testing.T) {
	type fields struct {
		endpoint endpoints.IEndpoint
	}
	tests := []struct {
		name        string
		fields      fields
		expectedErr error
	}{
		{
			name: "given_endpoint_fail_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{}, errors.New("mock_error"))
					return endpoint
				}(),
			},
			expectedErr: errDoRequest,
		},
		{
			name: "given_endpoint_status_conflict_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 409,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: ErrAccountConflict,
		},
		{
			name: "given_endpoint_status_no_content_then_return_ok",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 204,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointDeleteAccount: tt.fields.endpoint,
				},
			}
			err := client.requestDeleteAccount("any_id", 0)
			require.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func Test_accountClient_requestFetchAccount(t *testing.T) {
	type fields struct {
		endpoint endpoints.IEndpoint
	}
	tests := []struct {
		name        string
		fields      fields
		expectedOut []byte
		expectedErr error
	}{
		{
			name: "given_endpoint_fail_request_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{}, errors.New("mock_error"))
					return endpoint
				}(),
			},
			expectedErr: errDoRequest,
		},
		{
			name: "given_endpoint_status_not_found_then_return_error",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 404,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: ErrAccountNotFound,
		},
		{
			name: "given_endpoint_status_ok_then_return_ok",
			fields: fields{
				endpoint: func() endpoints.IEndpoint {
					endpoint := &endpointMock{}
					endpoint.On("Do", mock.Anything).
						Return(&http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader("{}")),
						}, nil)
					return endpoint
				}(),
			},
			expectedErr: nil,
			expectedOut: []byte("{}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointFetchAccount: tt.fields.endpoint,
				},
			}
			got, err := client.requestFetchAccount("any_id")
			require.True(t, errors.Is(err, tt.expectedErr))
			require.Equal(t, tt.expectedOut, got)
		})
	}
}
