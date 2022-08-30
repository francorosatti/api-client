package account_test

import (
	"github.com/google/uuid"
	"testing"

	"github.com/francorosatti/form3-api-client/pkg/form3"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
	"github.com/stretchr/testify/assert"
)

func Test_Integration_CreateAccount(t *testing.T) {
	// Arrange
	client, err := form3.NewClient(form3.EnvironmentLocal)
	assert.NoError(t, err)

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
		WithAttributes(*accountAttributes)

	account := models.Account{}
	account.WithData(accountData)

	uuid.New()
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

	accountData := models.AccountData{}
	accountData.WithID("id").
		WithOrganisationID("organisation_id")

	account := models.Account{}
	account.WithData(accountData)

	createdAccount, err := client.CreateAccount(account)

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

	accountData := models.AccountData{}
	accountData.WithID("id").
		WithOrganisationID("organisation_id")

	account := models.Account{}
	account.WithData(accountData)

	createdAccount, err := client.CreateAccount(account)

	// Act
	err = client.DeleteAccount(createdAccount.Data.ID)

	// Assert
	assert.NoError(t, err)

	fetchedAccount, err := client.FetchAccount(createdAccount.Data.ID)
	assert.Nil(t, fetchedAccount)
	assert.Error(t, err)
	// assert.True(t, errors.Is(err, accounts.ErrAccountNotFound))
}
