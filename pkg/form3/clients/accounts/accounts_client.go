package accounts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/francorosatti/form3-api-client/pkg/form3/internal/endpoints"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

type IAccountClient interface {
	CreateAccount(models.Account) (models.Account, error)
	FetchAccount(accountID string) (models.Account, error)
	DeleteAccount(accountID string, version int64) error
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
			&http.Client{Timeout: 3 * time.Second},
			fmt.Sprintf("%s/organisation/accounts", baseUrl),
			http.MethodPost,
		),
		_endpointFetchAccount: endpoints.NewEndpoint(
			&http.Client{Timeout: 3 * time.Second},
			fmt.Sprintf("%s/organisation/accounts/{id}", baseUrl),
			http.MethodGet,
		),
		_endpointDeleteAccount: endpoints.NewEndpoint(
			&http.Client{Timeout: 3 * time.Second},
			fmt.Sprintf("%s/organisation/accounts/{id}", baseUrl),
			http.MethodDelete,
		),
	}
}

func (client accountClient) CreateAccount(account models.Account) (models.Account, error) {
	accountBytes, err := accountDataToJson(account)
	if err != nil {
		return models.Account{}, err
	}

	response, err := client.requestCreateAccount(accountBytes)
	if err != nil {
		return models.Account{}, err
	}

	return jsonToAccountData(response)
}

func (client accountClient) FetchAccount(accountID string) (models.Account, error) {
	response, err := client.requestFetchAccount(accountID)
	if err != nil {
		return models.Account{}, err
	}

	return jsonToAccountData(response)
}

func (client accountClient) DeleteAccount(accountID string, version int64) error {
	return client.requestDeleteAccount(accountID, version)
}

func (client accountClient) requestCreateAccount(accountBody []byte) ([]byte, error) {
	endpoint := client.endpoints[_endpointCreateAccount]

	requestBody := endpoints.WithBody(accountBody)

	res, err := endpoint.Do(requestBody)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errResponseReadBody, err)
	}

	if err = handleStatusCode(res.StatusCode, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (client accountClient) requestFetchAccount(id string) ([]byte, error) {
	endpoint := client.endpoints[_endpointFetchAccount]

	params := endpoints.WithParam(_paramID, id)

	res, err := endpoint.Do(params)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errResponseReadBody, err)
	}

	if err = handleStatusCode(res.StatusCode, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (client accountClient) requestDeleteAccount(id string, version int64) error {
	endpoint := client.endpoints[_endpointDeleteAccount]

	params := endpoints.WithParam(_paramID, id)
	query := endpoints.WithQueryParam(_queryVersion, fmt.Sprintf("%d", version))

	res, err := endpoint.Do(params, query)
	if err != nil {
		return fmt.Errorf("%w: %s", errDoRequest, err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%w: %s", errResponseReadBody, err)
	}

	return handleStatusCode(res.StatusCode, body)
}

func handleStatusCode(statusCode int, body []byte) error {
	if statusCode < 300 {
		return nil
	}

	switch statusCode {
	case 400:
		return fmt.Errorf("%w: %s", ErrAccountBadRequest, string(body))
	case 404:
		return ErrAccountNotFound
	case 409:
		return ErrAccountConflict
	default:
		return fmt.Errorf("%w: status code %d", errResponseStatusCode, statusCode)
	}
}
