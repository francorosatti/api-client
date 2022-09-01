package account_test

import (
	"testing"

	"github.com/francorosatti/form3-api-client/pkg/form3"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Integration_CreateAccount(t *testing.T) {
	// Arrange
	client, err := form3.NewClient(form3.EnvironmentLocal)
	assert.NoError(t, err)

	account := getTestAccount()

	// Act
	createdAccount, err := client.CreateAccount(account)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, account, createdAccount)
}

func Test_Integration_FetchAccount(t *testing.T) {
	// Arrange
	client, err := form3.NewClient(form3.EnvironmentLocal)
	assert.NoError(t, err)

	account := getTestAccount()

	createdAccount, err := client.CreateAccount(account)
	require.NoError(t, err)

	// Act
	fetchedAccount, err := client.FetchAccount(createdAccount.Data.ID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createdAccount, fetchedAccount)
}

func Test_Integration_DeleteAccount(t *testing.T) {
	// Arrange
	client, err := form3.NewClient(form3.EnvironmentLocal)
	assert.NoError(t, err)

	account := getTestAccount()

	createdAccount, err := client.CreateAccount(account)

	// Act
	err = client.DeleteAccount(createdAccount.Data.ID, *createdAccount.Data.Version)

	// Assert
	assert.NoError(t, err)

	fetchedAccount, err := client.FetchAccount(createdAccount.Data.ID)
	assert.Equal(t, models.Account{}, fetchedAccount)
	assert.Error(t, err)
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
