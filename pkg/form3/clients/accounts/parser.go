package accounts

import (
	"encoding/json"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

func serialiseAccountData(model models.Account) ([]byte, error) {
	return json.Marshal(model)
}

func parseAccountData(bytes []byte) (models.Account, error) {
	var model models.Account
	err := json.Unmarshal(bytes, &model)
	return model, err
}
