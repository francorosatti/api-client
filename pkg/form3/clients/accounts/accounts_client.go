package accounts

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/francorosatti/form3-api-client/pkg/form3/internal/endpoints"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

type IAccountClient interface {
	CreateAccount(models.Account) (models.Account, error)
	FetchAccount(accountID string) (models.Account, error)
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
			fmt.Sprintf("%s/organisation/accounts", baseUrl),
			http.MethodPost,
		),
		_endpointFetchAccount: endpoints.NewEndpoint(
			&http.Client{},
			fmt.Sprintf("%s/organisation/accounts/{id}", baseUrl),
			http.MethodGet,
		),
		_endpointDeleteAccount: endpoints.NewEndpoint(
			&http.Client{},
			fmt.Sprintf("%s/organisation/accounts/{id}", baseUrl),
			http.MethodDelete,
		),
	}
}

func (client accountClient) CreateAccount(account models.Account) (models.Account, error) {
	accountBytes, err := serialiseAccountData(account)
	if err != nil {
		return models.Account{}, err
	}

	response, err := client.createAccountRequest(accountBytes)
	if err != nil {
		return models.Account{}, err
	}

	return parseAccountData(response)
}

func (client accountClient) FetchAccount(accountID string) (models.Account, error) {
	response, err := client.fetchAccountRequest(accountID)
	if err != nil {
		return models.Account{}, err
	}

	return parseAccountData(response)
}

func (client accountClient) DeleteAccount(accountID string) error {
	return client.deleteAccountRequest(accountID)
}

func (client accountClient) createAccountRequest(accountBody []byte) ([]byte, error) {
	endpoint := client.endpoints[_endpointCreateAccount]

	requestBody := endpoints.WithBody(accountBody)

	res, err := endpoint.Do(requestBody)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%w: status code %d", errResponseStatusCode, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errResponseReadBody, err)
	}

	return body, nil
}

func (client accountClient) fetchAccountRequest(id string) ([]byte, error) {
	endpoint := client.endpoints[_endpointFetchAccount]

	params := endpoints.WithParam(_paramID, id)

	res, err := endpoint.Do(params)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%w: status code %d", errResponseStatusCode, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errResponseReadBody, err)
	}

	return body, nil
}

func (client accountClient) deleteAccountRequest(id string) error {
	endpoint := client.endpoints[_endpointDeleteAccount]

	params := endpoints.WithParam(_paramID, id)

	res, err := endpoint.Do(params)
	if err != nil {
		return fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("%w: status code %d", errResponseStatusCode, res.StatusCode)
	}

	return nil
}
