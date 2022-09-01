package accounts

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/francorosatti/form3-api-client/pkg/form3/internal/endpoints"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			name: "given any input" +
				"when endpoint responds status bad request" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds invalid json" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status ok" +
				"then return response with model",
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
			// Act
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointCreateAccount: tt.fields.endpoint,
				},
			}

			// Act
			got, err := client.CreateAccount(models.Account{})

			// Assert
			assert.True(t, errors.Is(err, tt.expectedErr))
			assert.Equal(t, tt.expectedOut, got)
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
			name: "given any input" +
				"when endpoint responds status not found" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status ok" +
				"then return response with model",
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
			// Arrange
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointFetchAccount: tt.fields.endpoint,
				},
			}

			// Act
			got, err := client.FetchAccount("any_id")

			//Assert
			assert.True(t, errors.Is(err, tt.expectedErr))
			assert.Equal(t, tt.expectedOut, got)
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
			name: "given any input" +
				"when endpoint request fails" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status bad request" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status ok" +
				"then return ok",
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
			// Arrange
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointCreateAccount: tt.fields.endpoint,
				},
			}

			// Act
			got, err := client.requestCreateAccount([]byte("anything"))

			// Assert
			assert.True(t, errors.Is(err, tt.expectedErr))
			assert.Equal(t, tt.expectedOut, got)
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
			name: "given any input" +
				"when endpoint request fails" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status conflict" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status no content" +
				"then return ok",
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
			// Act
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointDeleteAccount: tt.fields.endpoint,
				},
			}

			// Act
			err := client.requestDeleteAccount("any_id", 0)

			// Assert
			assert.True(t, errors.Is(err, tt.expectedErr))
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
			name: "given any input" +
				"when endpoint request fails" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status not found" +
				"then return error",
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
			name: "given any input" +
				"when endpoint responds status ok" +
				"then return ok",
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
			// Arrange
			client := accountClient{
				endpoints: map[string]endpoints.IEndpoint{
					_endpointFetchAccount: tt.fields.endpoint,
				},
			}

			// Act
			got, err := client.requestFetchAccount("any_id")

			// Assert
			assert.True(t, errors.Is(err, tt.expectedErr))
			assert.Equal(t, tt.expectedOut, got)
		})
	}
}
