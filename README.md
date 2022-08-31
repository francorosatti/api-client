# form3-api-client
Client for Form3 fake API

## Changelog
Any changes to this project can be found [here](./CHANGELOG.md)

## Integration Tests
In order to run integrations tests follow these steps:
1. Go to project root directory
2. Set up integration tests environment
    ```sh
    docker-compose up
    ```
3. Run integration tests
    ```sh
    go test .\pkg\integration_test\...
    ```

## Unitary Tests
In order to run unitary tests follow these steps:
1. Go to project root directory
2. Run unitary tests
    ```sh
    go test .\pkg\form3\...
    ```

## Client Usage
First import this library from your project.
```sh
go got github.com/francorosatti/form3-api-client
```

Then create a new client.
```go
import "github.com/francorosatti/form3-api-client/pkg/form3"

client, err := form3.NewClient(form3.EnvironmentLocal)
```

Finally, call the needed services.

```go
account := models.Account{...}
createdAccount, err := client.CreateAccount(account)
```

## Account Services
Available services for accounts are defined in [this interface](./pkg/form3/clients/accounts/accounts_client.go):
```go
type IAccountClient interface {
	CreateAccount(models.Account) (models.Account, error)
	FetchAccount(accountID string) (models.Account, error)
	DeleteAccount(accountID string, version int64) error
}
```

Models can be found [here](./pkg/form3/models)