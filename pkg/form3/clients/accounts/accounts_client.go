package accounts

import (
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

type IAccountClient interface {
	CreateAccount(models.AccountData) (models.AccountData, error)
	FetchAccount(accountID string) (models.AccountData, error)
	DeleteAccount(accountID string) error
}

type accountClient struct {
	baseUrl string
}

func NewAccountClient(baseUrl string) IAccountClient {
	return accountClient{
		baseUrl: baseUrl,
	}
}

func (client accountClient) CreateAccount(account models.AccountData) (models.AccountData, error) {
	panic("not_implemented")
}

func (client accountClient) FetchAccount(accountID string) (models.AccountData, error) {
	panic("not_implemented")
}

func (client accountClient) DeleteAccount(accountID string) error {
	panic("not_implemented")
}
