package form3

import (
	"fmt"
	"github.com/francorosatti/form3-api-client/pkg/form3/clients/accounts"
)

type IClient interface {
	accounts.IAccountClient
}

type client struct {
	accounts.IAccountClient
}

func NewClient(env Environment) (IClient, error) {
	host, exists := _hostByEnvironment[env]
	if !exists {
		return nil, ErrUnknownEnvironment
	}

	baseUrl := fmt.Sprintf("%s/%s", host, _apiVersion)

	return client{
		IAccountClient: accounts.NewAccountClient(baseUrl),
	}, nil
}
