package accounts

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/francorosatti/form3-api-client/pkg/form3/internal/endpoints"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

type IAccountClient interface {
	CreateAccount(models.AccountData) (models.AccountData, error)
	FetchAccount(accountID string) (models.AccountData, error)
	DeleteAccount(accountID string) error
}

type accountClient struct {
	endpoints map[string]endpoints.IEndpoint
}

func NewAccountClient(baseUrl string) IAccountClient {
	endpoints := createEndpoints(baseUrl)

	return accountClient{
		endpoints: endpoints,
	}
}

func createEndpoints(baseUrl string) map[string]endpoints.IEndpoint {
	return map[string]endpoints.IEndpoint{
		_endpointCreateAccount: endpoints.NewEndpoint(
			&http.Client{},
			fmt.Sprintf("%s/accounts", baseUrl),
			http.MethodPost,
		),
		_endpointGetAccount: endpoints.NewEndpoint(
			&http.Client{},
			fmt.Sprintf("%s/accounts/{id}", baseUrl),
			http.MethodGet,
		),
		_endpointDeleteAccount: endpoints.NewEndpoint(
			&http.Client{},
			fmt.Sprintf("%s/accounts/{id}", baseUrl),
			http.MethodDelete,
		),
	}
}

func (client accountClient) CreateAccount(account models.AccountData) (models.AccountData, error) {
	accountBytes, err := serialiseAccountData(account)
	if err != nil {
		return models.AccountData{}, err
	}

	response, err := client.createAccountRequest(accountBytes)
	if err != nil {
		return models.AccountData{}, err
	}

	return parseAccountData(response)
}

func (client accountClient) FetchAccount(accountID string) (models.AccountData, error) {
	response, err := client.fetchAccountRequest(accountID)
	if err != nil {
		return models.AccountData{}, err
	}

	return parseAccountData(response)
}

func (client accountClient) DeleteAccount(accountID string) error {
	return client.deleteAccountRequest(accountID)
}

func (client accountClient) createAccountRequest(accountBody []byte) ([]byte, error) {
	return []byte{}, nil
}

func (client accountClient) fetchAccountRequest(id string) ([]byte, error) {
	endpoint := client.endpoints[_endpointGetAccount]

	params := endpoints.WithParam(_paramID, id)

	res, err := endpoint.Do(params)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request error: code %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (client accountClient) deleteAccountRequest(id string) error {
	return nil
}
