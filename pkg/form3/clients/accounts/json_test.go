package accounts

import (
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseAccountData(t *testing.T) {
	tests := []struct {
		name    string
		json    []byte
		want    models.Account
		wantErr bool
	}{
		{
			name: "given_valid_json_then_return_object",
			json: []byte(`{"data":{"id":"id","attributes":{"bank_id":"bank_id"}}}`),
			want: models.Account{
				Data: &models.AccountData{
					ID: "id",
					Attributes: &models.AccountAttributes{
						BankID: "bank_id",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "given_invalid_json_then_return_error",
			json:    []byte(`}`),
			want:    models.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonToAccountData(tt.json)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_serialiseAccountData(t *testing.T) {
	model := models.Account{
		Data: &models.AccountData{
			ID: "id",
			Attributes: &models.AccountAttributes{
				BankID: "bank_id",
			},
		},
	}

	expectedJson := []byte(`{"data":{"attributes":{"bank_id":"bank_id"},"id":"id"}}`)

	got, err := accountDataToJson(model)
	assert.NoError(t, err)
	assert.Equal(t, expectedJson, got)
}
