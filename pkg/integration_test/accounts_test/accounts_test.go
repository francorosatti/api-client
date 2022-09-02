package account_test

import (
	"errors"
	"testing"

	"github.com/francorosatti/form3-api-client/pkg/form3"
	"github.com/francorosatti/form3-api-client/pkg/form3/clients/accounts"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Integration_CreateAccount(t *testing.T) {
	type args struct {
		model *models.Account
	}
	tests := []struct {
		name        string
		args        args
		expectedOut models.Account
		expectedErr error
	}{
		{
			name: "given account with all valid fields" +
				"when creating account" +
				"then return created account",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID("9ab01f5b-c1e3-4596-8f3b-4e98b8f57785").
						WithOrganisationID("200231e0-f512-4d95-93db-934820c0a156").
						WithType("accounts").
						WithVersion(0).
						WithAttributes(
							*new(models.AccountAttributes).
								WithAccountClassification("Personal").
								WithAccountMatchingOptOut(true).
								WithAccountNumber("12345678").
								WithAlternativeNames([]string{"alternative_names"}).
								WithBaseCurrency("GBP").
								WithBankID("ABCD").
								WithBankIDCode("ABCDEF").
								WithBic("NWBKGB22").
								WithCountry("GB").
								WithIban("AA00").
								WithJointAccount(true).
								WithName([]string{"account_name"}).
								WithSecondaryIdentification("secondary_identification").
								WithStatus("confirmed").
								WithSwitched(true),
						),
				),
			},
			expectedOut: *new(models.Account).WithData(
				*new(models.AccountData).
					WithID("9ab01f5b-c1e3-4596-8f3b-4e98b8f57785").
					WithOrganisationID("200231e0-f512-4d95-93db-934820c0a156").
					WithType("accounts").
					WithVersion(0).
					WithAttributes(
						*new(models.AccountAttributes).
							WithAccountClassification("Personal").
							WithAccountMatchingOptOut(true).
							WithAccountNumber("12345678").
							WithAlternativeNames([]string{"alternative_names"}).
							WithBaseCurrency("GBP").
							WithBankID("ABCD").
							WithBankIDCode("ABCDEF").
							WithBic("NWBKGB22").
							WithCountry("GB").
							WithIban("AA00").
							WithJointAccount(true).
							WithName([]string{"account_name"}).
							WithSecondaryIdentification("secondary_identification").
							WithStatus("confirmed").
							WithSwitched(true),
					),
			),
			expectedErr: nil,
		},
		{
			name: "given account with the minimum valid fields" +
				"when creating account" +
				"then return created account",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID("e414146c-5e9f-4c68-942c-6573d7d14d90").
						WithOrganisationID("a3b33ec8-9ee4-42ec-b436-561f0264fc57").
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB").
								WithName([]string{"account_name"}),
						),
				),
			},
			expectedOut: *new(models.Account).WithData(
				*new(models.AccountData).
					WithID("e414146c-5e9f-4c68-942c-6573d7d14d90").
					WithOrganisationID("a3b33ec8-9ee4-42ec-b436-561f0264fc57").
					WithType("accounts").
					WithVersion(0).
					WithAttributes(
						*new(models.AccountAttributes).
							WithCountry("GB").
							WithName([]string{"account_name"}),
					),
			),
			expectedErr: nil,
		},
		{
			name: "given account with invalid account id" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID("invalid_id").
						WithOrganisationID(uuid.NewString()).
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB"),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given account with invalid organisation id" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID(uuid.NewString()).
						WithOrganisationID("invalid_organisation_id").
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB"),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given account without name" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID(uuid.NewString()).
						WithOrganisationID(uuid.NewString()).
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB"),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given account without country" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID(uuid.NewString()).
						WithOrganisationID(uuid.NewString()).
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithName([]string{"account_name"}),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given account with invalid iban" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID(uuid.NewString()).
						WithOrganisationID(uuid.NewString()).
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB").
								WithIban("invalid_iban").
								WithName([]string{"account_name"}),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given account with invalid bank_id" +
				"when creating account" +
				"then return error",
			args: args{
				model: new(models.Account).WithData(
					*new(models.AccountData).
						WithID(uuid.NewString()).
						WithOrganisationID(uuid.NewString()).
						WithType("accounts").
						WithAttributes(
							*new(models.AccountAttributes).
								WithCountry("GB").
								WithBankID("invalid_bank_id").
								WithName([]string{"account_name"}),
						),
				),
			},
			expectedOut: models.Account{},
			expectedErr: accounts.ErrAccountBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			client, err := form3.NewClient(form3.EnvironmentLocal)
			require.NoError(t, err)

			// Act
			got, err := client.CreateAccount(*tt.args.model)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func Test_Integration_FetchAccount(t *testing.T) {
	client, err := form3.NewClient(form3.EnvironmentLocal)
	require.NoError(t, err)

	testAccount := getTestAccount()
	createdAccount, err := client.CreateAccount(testAccount)
	require.NoError(t, err)

	type args struct {
		accountID string
	}
	tests := []struct {
		name        string
		args        args
		expectedOut models.Account
		expectedErr error
	}{
		{
			name: "given valid account id" +
				"when fetching account" +
				"then return account",
			args: args{
				accountID: createdAccount.Data.ID,
			},
			expectedOut: createdAccount,
			expectedErr: nil,
		},
		{
			name: "given invalid account id" +
				"when fetching account" +
				"then return not found",
			args: args{
				accountID: "invalid_id",
			},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given empty account id" +
				"when fetching account" +
				"then return not found",
			args: args{
				accountID: "",
			},
			expectedErr: accounts.ErrAccountInvalidParameters,
		},
		{
			name: "given non existent account id" +
				"when fetching account" +
				"then return not found",
			args: args{
				accountID: uuid.NewString(),
			},
			expectedErr: accounts.ErrAccountNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := client.FetchAccount(tt.args.accountID)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func Test_Integration_DeleteAccount(t *testing.T) {
	client, err := form3.NewClient(form3.EnvironmentLocal)
	require.NoError(t, err)

	testAccount := getTestAccount()
	createdAccount, err := client.CreateAccount(testAccount)
	require.NoError(t, err)

	type args struct {
		accountID string
	}
	tests := []struct {
		name        string
		args        args
		expectedOut models.Account
		expectedErr error
	}{
		{
			name: "given valid account id" +
				"when deleting account" +
				"then return ok",
			args: args{
				accountID: createdAccount.Data.ID,
			},
			expectedOut: createdAccount,
			expectedErr: nil,
		},
		{
			name: "given invalid account id" +
				"when deleting account" +
				"then return not found",
			args: args{
				accountID: "invalid_id",
			},
			expectedErr: accounts.ErrAccountBadRequest,
		},
		{
			name: "given empty account id" +
				"when deleting account" +
				"then return not found",
			args: args{
				accountID: "",
			},
			expectedErr: accounts.ErrAccountInvalidParameters,
		},
		{
			name: "given non existent account id" +
				"when deleting account" +
				"then return not found",
			args: args{
				accountID: uuid.NewString(),
			},
			expectedErr: accounts.ErrAccountNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := client.FetchAccount(tt.args.accountID)

			// Assert
			assert.Equal(t, tt.expectedOut, got)
			assert.True(t, errors.Is(err, tt.expectedErr))
		})
	}
}

func getTestAccount() models.Account {
	accountAttributes := &models.AccountAttributes{}
	accountAttributes.WithCountry("GB").
		WithBankID("123456").
		WithBic("NWBKGB22").
		WithBankIDCode("GBDSC").
		WithBaseCurrency("GBP").
		WithName([]string{"account name"})

	accountData := models.AccountData{}
	accountData.WithID(uuid.NewString()).
		WithOrganisationID(uuid.NewString()).
		WithType("accounts").
		WithVersion(0).
		WithAttributes(*accountAttributes)

	account := models.Account{}
	account.WithData(accountData)

	return account
}
